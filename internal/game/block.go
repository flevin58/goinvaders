package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Block struct {
	position rl.Vector2
	active   bool
}

func NewBlock(x float32, y float32) *Block {
	return &Block{
		position: rl.Vector2{X: x, Y: y},
		active:   true,
	}
}

func (b *Block) GetRect() rl.Rectangle {
	return rl.Rectangle{
		X:      b.position.X,
		Y:      b.position.Y,
		Width:  3,
		Height: 3,
	}
}

func (b *Block) Draw() {
	blockColor := rl.Color{R: 243, G: 216, B: 63, A: 255}
	rl.DrawRectangle(int32(b.position.X), int32(b.position.Y), 3, 3, blockColor)
}
