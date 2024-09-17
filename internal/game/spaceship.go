package game

import (
	"goinvaders/internal/assets"
	"goinvaders/internal/tools"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Spaceship struct {
	image        rl.Texture2D
	position     rl.Vector2
	lasers       []*Laser
	lastFireTime float64
	laserSound   rl.Sound
	mute         bool
}

func NewSpaceship() Spaceship {
	image := assets.GetSpaceshipImage()
	xpos := float32(rl.GetScreenWidth()-int(image.Width)) / 2
	ypos := float32(rl.GetScreenHeight()-int(image.Height)) - 100
	return Spaceship{
		image:        image,
		position:     rl.Vector2{X: xpos, Y: ypos},
		lasers:       make([]*Laser, 0),
		lastFireTime: 0,
		laserSound:   assets.LoadSound("laser.ogg"),
		mute:         false,
	}
}

func (s *Spaceship) GetRect() rl.Rectangle {
	return rl.Rectangle{
		X:      s.position.X,
		Y:      s.position.Y,
		Width:  float32(s.image.Width),
		Height: float32(s.image.Height),
	}
}

func (s *Spaceship) Reset() {
	s.position.X = float32(rl.GetScreenWidth()-int(s.image.Width)) / 2
	s.position.Y = float32(rl.GetScreenHeight()) - float32(s.image.Height) - 100
	s.lasers = make([]*Laser, 0)
}

func (s *Spaceship) FireLaser() {
	if rl.GetTime()-s.lastFireTime >= 0.35 {
		if !s.mute {
			rl.PlaySound(s.laserSound)
		}
		posx := int32(s.position.X) + s.image.Width/2 - 2
		posy := int32(s.position.Y)
		s.lasers = append(s.lasers, NewLaser(posx, posy, -6))
		s.lastFireTime = rl.GetTime()
	}
}

func (s *Spaceship) Update() {
	// delete inactive lasers
	s.lasers = tools.FilterSlice(s.lasers,
		func(laser *Laser) bool {
			return laser.active
		})

	for _, laser := range s.lasers {
		laser.Update()
	}
}

func (s *Spaceship) Draw() {
	rl.DrawTextureV(s.image, s.position, rl.White)

	for _, laser := range s.lasers {
		laser.Draw()
	}
}

func (s *Spaceship) DrawAt(xpos, ypos int32) {
	position := rl.Vector2{X: float32(xpos), Y: float32(ypos)}
	rl.DrawTextureV(s.image, position, rl.White)
}

func (s *Spaceship) MoveLeft() {
	s.position.X -= 7
	if s.position.X < 25 {
		s.position.X = 25
	}
}

func (s *Spaceship) MoveRight() {
	s.position.X += 7
	maxpos := float32(rl.GetScreenWidth()-int(s.image.Width)) - 25
	if s.position.X > maxpos {
		s.position.X = maxpos
	}
}
