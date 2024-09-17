package game

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Laser struct {
	position rl.Vector2
	speed    float32
	active   bool
}

func NewLaser(posx, posy int32, speed float32) *Laser {
	return &Laser{
		position: rl.Vector2{X: float32(posx), Y: float32(posy)},
		speed:    speed,
		active:   true,
	}
}

func (l *Laser) GetRect() rl.Rectangle {
	return rl.Rectangle{
		X:      l.position.X,
		Y:      l.position.Y,
		Width:  4,
		Height: 15,
	}
}

func (l *Laser) CollidedWith(other Collideable) bool {
	return rl.CheckCollisionRecs(other.GetRect(), l.GetRect())
}

func (l *Laser) IsActive() bool {
	return l.active
}

func (l *Laser) Update() {
	if l.active {
		l.position.Y += l.speed
		if (l.position.Y > float32(rl.GetScreenHeight())-100) || (l.position.Y < 25) {
			l.active = false
		}
	}
}

func (l *Laser) Draw() {
	if l.active {
		rl.DrawRectangle(int32(l.position.X), int32(l.position.Y), 4, 15, color.RGBA{R: 243, G: 216, B: 63, A: 255})
	}
}
