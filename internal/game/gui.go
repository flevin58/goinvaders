package game

import (
	"fmt"
	"goinvaders/internal/assets"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) TextAt(posx int, posy int, text string, args ...any) {
	if len(args) > 0 {
		text = fmt.Sprintf(text, args...)
	}
	rl.DrawTextEx(g.font, text, rl.Vector2{X: float32(posx), Y: float32(posy)}, 34, 2, assets.Yellow)
}

func (g *Game) CenterTextAt(posx int, posy int, width int, text string, args ...any) {
	if len(args) > 0 {
		text = fmt.Sprintf(text, args...)
	}
	textWidth := int(rl.MeasureTextEx(g.font, text, 34, 2).X)
	posx += (width - textWidth) / 2
	rl.DrawTextEx(g.font, text, rl.Vector2{X: float32(posx), Y: float32(posy)}, 34, 2, assets.Yellow)
}

func (g *Game) DrawDialogBox(text1, text2, text3 string, bkgcolor rl.Color) {
	rwidth := 500
	rheight := 200
	rposx := (rl.GetScreenWidth() - rwidth) / 2
	rposy := 100

	rec := rl.Rectangle{
		X:      float32(rposx),
		Y:      float32(rposy),
		Width:  float32(rwidth),
		Height: float32(rheight),
	}

	rl.DrawRectangleGradientH(int32(rposx), int32(rposy), int32(rwidth), int32(rheight), bkgcolor, bkgcolor)
	rl.DrawRectangleLinesEx(rec, 10.0, yellow)
	g.CenterTextAt(rposx, 150, rwidth, text1)
	g.CenterTextAt(rposx, 190, rwidth, text2)
	g.CenterTextAt(rposx, 230, rwidth, text3)
}

func (g *Game) GameOverDraw() {
	g.DrawDialogBox("GAME OVER", "PRESS ENTER TO RESTART", "PRESS ESC TO QUIT", red)
}

func (g *Game) LevelUpDraw() {
	g.DrawDialogBox("CONGRATULATIONS", "YOU DEFEATED THE ALIENS", "PRESS ENTER FOR NEXT LEVEL", green)
}
