package game

import (
	"goinvaders/internal/assets"
	"goinvaders/internal/tools"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Collideable interface {
	GetRect() rl.Rectangle
}

const alienLaserShootInterval float64 = 0.35

var (
	grey   = color.RGBA{R: 29, G: 29, B: 27, A: 255}
	yellow = color.RGBA{R: 243, G: 216, B: 63, A: 255}
)

type Game struct {
	spaceship          Spaceship
	mysteryship        MysteryShip
	obstacles          []*Obstacle
	aliens             []*Alien
	aliensDirection    int32
	alienLasers        []*Laser
	timeLastAlienFired float64
	msSpawnInterval    float64
	msTimeLastSpawned  float64
	lives              int32
	running            bool
	gameover           bool
	quit               bool
	font               rl.Font
	level              int32
	score              int32
	highScore          int32
	music              rl.Music
	explosionSound     rl.Sound
	mutesfx            bool
	mutemusic          bool
}

func New() Game {
	game := Game{
		spaceship:      NewSpaceship(),
		mysteryship:    NewMysteryShip(),
		font:           assets.LoadFont("monogram.ttf"),
		music:          assets.LoadMusic("music.ogg"),
		explosionSound: assets.LoadSound("explosion.ogg"),
		mutesfx:        false,
		mutemusic:      false,
	}

	game.InitGame()
	if !rl.IsMusicReady(game.music) {
		rl.TraceLog(rl.LogError, "Music not ready")
	}
	rl.PlayMusicStream(game.music)
	rl.SetMusicVolume(game.music, 0.6)
	return game
}

func (g *Game) InitGame() {
	g.aliensDirection = 1
	g.running = true
	g.quit = false
	g.gameover = false
	g.lives = 3
	g.msSpawnInterval = float64(rl.GetRandomValue(10, 20))
	g.msTimeLastSpawned = 0
	g.timeLastAlienFired = 0
	g.level = 1
	g.score = 0
	g.highScore = 0
	g.LoadHighScore()
	g.CreateObstacles()
	g.CreateAliens()
}

func (g *Game) ResetGame() {
	g.spaceship.Reset()
	g.aliens = make([]*Alien, 0)
	g.alienLasers = make([]*Laser, 0)
	g.obstacles = make([]*Obstacle, 0)
}

func (g *Game) CreateObstacles() {
	obstacleWidth := GetObstacleWidth()
	gap := (rl.GetScreenWidth() - (4 * obstacleWidth)) / 5
	for i := range 4 {
		offsetx := (i+1)*gap + i*obstacleWidth
		g.obstacles = append(g.obstacles, NewObstacle(float32(offsetx), float32(rl.GetScreenHeight()-200)))
	}
}

func (g *Game) CreateAliens() {
	for row := range 5 {
		var alienType int32
		switch {
		case row == 0:
			alienType = 3
		case row == 1 || row == 2:
			alienType = 2
		default:
			alienType = 1
		}
		for col := range 11 {
			posx := 75 + col*55
			posy := 110 + row*55
			g.aliens = append(g.aliens, NewAlien(alienType, int32(posx), int32(posy)))
		}
	}
}

func (g *Game) MoveDownAliens(distance int) {
	for _, alien := range g.aliens {
		alien.position.Y += float32(distance)
	}
}

func (g *Game) MoveAliens() {
	for _, alien := range g.aliens {
		if alien.position.X+float32(alien.image.Width) > float32(rl.GetScreenWidth())-25 {
			g.aliensDirection = -1
			g.MoveDownAliens(4)
		}
		if alien.position.X < 25 {
			g.aliensDirection = 1
			g.MoveDownAliens(4)
		}
		alien.Update(g.aliensDirection)
	}
}

func (g *Game) AliensShootLaser() {

	// there must be an alien
	if len(g.aliens) <= 0 {
		return
	}

	// enough time should have passed from last alien laser
	if rl.GetTime()-g.timeLastAlienFired < alienLaserShootInterval {
		return
	}

	// create a random alien laser and add it to the queue
	randomIndex := rl.GetRandomValue(0, int32(len(g.aliens)-1))
	alien := g.aliens[randomIndex]
	laserx := int32(alien.position.X) + alien.image.Width/2
	lasery := int32(alien.position.Y) + alien.image.Height
	g.alienLasers = append(g.alienLasers, NewLaser(laserx, lasery, 6))
	g.timeLastAlienFired = rl.GetTime()
}

func (g *Game) AddScore(earned int32) {
	g.score += earned
	if g.score > g.highScore {
		g.highScore = g.score
	}
}

func (g *Game) CheckForCollisions() {
	// Spaceship lasers
	for _, laser := range g.spaceship.lasers {
		// Check against aliens
		deleteAliens := false
		for _, alien := range g.aliens {
			if laser.CollidedWith(alien) {
				if !g.mutesfx {
					rl.PlaySound(g.explosionSound)
				}
				g.AddScore(alien.GetScore())
				alien.active = false
				laser.active = false
				deleteAliens = true
			}
		}
		// If we deactivated some aliens, delete them
		if deleteAliens {
			g.aliens = tools.FilterSlice(g.aliens,
				func(alien *Alien) bool {
					return alien.active
				})
		}

		// Check against blocks
		for _, obstacle := range g.obstacles {
			deleteBlocks := false
			for _, block := range obstacle.blocks {
				if laser.CollidedWith(block) {
					block.active = false
					laser.active = false
					deleteBlocks = true
				}
			}
			if deleteBlocks {
				obstacle.blocks = tools.FilterSlice(obstacle.blocks,
					func(block *Block) bool {
						return block.active
					})
			}
		}

		// Check against mystery ship
		if laser.CollidedWith(&g.mysteryship) {
			if !g.mutesfx {
				rl.PlaySound(g.explosionSound)
			}
			g.AddScore(500)
			g.mysteryship.alive = false
			laser.active = false
		}
	}

	// Alien Lasers
	for _, laser := range g.alienLasers {
		// Alien lasers against Spaceship
		if laser.CollidedWith(&g.spaceship) {
			laser.active = false
			g.lives--
			if g.lives == 0 {
				g.GameOver()
			}
			rl.TraceLog(rl.LogInfo, "Spaceship hit")
		}
		// Alien lasers against Obstacles
		for _, obstacle := range g.obstacles {
			deleteBlocks := false
			for _, block := range obstacle.blocks {
				if laser.CollidedWith(block) {
					block.active = false
					laser.active = false
					deleteBlocks = true
				}
			}
			if deleteBlocks {
				obstacle.blocks = tools.FilterSlice(obstacle.blocks,
					func(block *Block) bool {
						return block.active
					})
			}
		}
	}

	for _, alien := range g.aliens {
		// Alien against obstacles
		for _, obstacle := range g.obstacles {
			deleteBlocks := false
			for _, block := range obstacle.blocks {
				if alien.CollidedWith(block) {
					block.active = false
					deleteBlocks = true
				}
			}
			if deleteBlocks {
				obstacle.blocks = tools.FilterSlice(obstacle.blocks,
					func(block *Block) bool {
						return block.active
					})
			}
		}
		// Alien against Spaceship
		if alien.CollidedWith(&g.spaceship) {
			g.GameOver()
		}
	}
}

func (g *Game) Update() {
	if !g.running {
		return
	}

	rl.UpdateMusicStream(g.music)

	g.CheckForCollisions()

	if rl.GetTime()-g.msTimeLastSpawned > g.msSpawnInterval {
		g.mysteryship.Spawn()
		g.msTimeLastSpawned = rl.GetTime()
		g.msSpawnInterval = float64(rl.GetRandomValue(10, 20))
	}
	g.spaceship.Update()
	g.mysteryship.Update()
	g.MoveAliens()

	// delete inactive lasers
	g.alienLasers = tools.FilterSlice(g.alienLasers,
		func(laser *Laser) bool {
			return laser.active
		})

	g.AliensShootLaser()
	for _, laser := range g.alienLasers {
		laser.Update()
	}
}

func (g *Game) Draw() {
	rl.ClearBackground(grey)

	// Draw the GUI
	rl.DrawRectangleRoundedLines(rl.Rectangle{X: 10, Y: 10, Width: 780, Height: 780}, 0.18, 20, 2, yellow)
	rl.DrawLineEx(rl.Vector2{X: 25, Y: 730}, rl.Vector2{X: 775, Y: 730}, 3, yellow)
	if g.running {
		g.TextAt(570, 740, "LEVEL %02d", g.level)
	} else {
		g.TextAt(570, 740, "GAME OVER")
	}
	for i := range g.lives {
		g.spaceship.DrawAt(50*(i+1), 745)
	}
	g.TextAt(50, 15, "SCORE")
	g.TextAt(50, 40, "%05d", g.score)

	g.TextAt(70, 15, "HIGH SCORE")
	g.TextAt(55, 40, "%05d", g.highScore)

	g.spaceship.Draw()
	g.mysteryship.Draw()

	for _, obstacle := range g.obstacles {
		obstacle.Draw()
	}

	for _, alien := range g.aliens {
		alien.Draw()
	}

	for _, laser := range g.alienLasers {
		laser.Draw()
	}

	if g.gameover {
		g.GameOverDraw()
	}
}

func (g *Game) ShouldQuit() bool {
	return g.quit || rl.WindowShouldClose()
}

func (g *Game) HandleInput() {
	if g.gameover {
		switch g.GameOverUpdate() {
		case Restart:
			g.ResetGame()
			g.InitGame()
		case End:
			g.quit = true
		}
	}

	if g.running {
		if rl.IsKeyDown(rl.KeyLeft) {
			g.spaceship.MoveLeft()
		} else if rl.IsKeyDown(rl.KeyRight) {
			g.spaceship.MoveRight()
		} else if rl.IsKeyDown(rl.KeySpace) {
			g.spaceship.FireLaser()
		}
	}

	// Pause/Resume music with key "M"
	if rl.IsKeyPressed(rl.KeyM) {
		g.mutemusic = !g.mutemusic
		if g.mutemusic {
			rl.PauseMusicStream(g.music)
		} else {
			rl.ResumeMusicStream(g.music)
		}
	}

	// Pause/Resume sfx with key "S"
	if rl.IsKeyPressed(rl.KeyS) {
		g.mutesfx = !g.mutesfx
		g.spaceship.mute = g.mutesfx
	}
}

func (g *Game) GameOver() {
	g.gameover = true
	g.running = false
	g.SaveHighScore()
	rl.TraceLog(rl.LogInfo, "Game Over!")
}
