package game

import (
	"goinvaders/internal/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MysteryShip struct {
	image    rl.Texture2D
	position rl.Vector2
	speed    int32
	alive    bool
}

func NewMysteryShip() MysteryShip {
	return MysteryShip{
		image: assets.GetMysteryImage(),
	}
}

func (m *MysteryShip) GetRect() rl.Rectangle {
	if m.alive {
		return rl.Rectangle{
			X:      m.position.X,
			Y:      m.position.Y,
			Width:  float32(m.image.Width),
			Height: float32(m.image.Height),
		}
	} else {
		return rl.Rectangle{
			X:      m.position.X,
			Y:      m.position.Y,
			Width:  0,
			Height: 0,
		}
	}
}

func (m *MysteryShip) Spawn() {
	m.position.Y = 90
	side := rl.GetRandomValue(0, 1)
	if side == 0 {
		m.position.X = 25
		m.speed = 3
	} else {
		m.position.X = float32(rl.GetScreenWidth()) - float32(m.image.Width) - 25
		m.speed = -3
	}
	m.alive = true
}

func (m *MysteryShip) Update() {
	if m.alive {
		m.position.X += float32(m.speed)
		if m.position.X > float32(rl.GetScreenWidth()-int(m.image.Width))-25 || m.position.X < 25 {
			m.alive = false
		}
	}
}

func (m *MysteryShip) Draw() {
	if m.alive {
		rl.DrawTextureV(m.image, m.position, rl.White)
	}
}
