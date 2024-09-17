package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Alien struct {
	alienType int32
	position  rl.Vector2
	image     rl.Texture2D
	active    bool
}

func NewAlien(image rl.Texture2D, alienType int32, xpos int32, ypos int32) *Alien {
	return &Alien{
		alienType: alienType,
		position:  rl.Vector2{X: float32(xpos), Y: float32(ypos)},
		image:     image,
		active:    true,
	}
}

func (a *Alien) GetRect() rl.Rectangle {
	return rl.Rectangle{
		X:      a.position.X,
		Y:      a.position.Y,
		Width:  float32(a.image.Width),
		Height: float32(a.image.Height),
	}
}

func (a *Alien) CollidedWith(other Collideable) bool {
	return rl.CheckCollisionRecs(other.GetRect(), a.GetRect())
}

func (a *Alien) GetScore() int32 {
	return a.alienType * 100
}

func (a *Alien) Update(direction int32) {
	a.position.X += float32(direction)
}

func (a *Alien) Draw() {
	rl.DrawTextureV(a.image, a.position, rl.White)
}
