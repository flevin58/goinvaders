package game

import (
	"fmt"
	"goinvaders/internal/tools"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Game) SaveHighScore() {
	fileName, err := tools.GetConfigPath("highscore.txt")
	if err != nil {
		rl.TraceLog(rl.LogError, err.Error())
	}
	file, err := os.Create(fileName)
	if err != nil {
		rl.TraceLog(rl.LogError, "Could not save high score to file: %s", fileName)
		return
	}
	fmt.Fprintf(file, "%d", g.highScore)
	file.Close()
}

func (g *Game) LoadHighScore() {
	fileName, err := tools.GetConfigPath("highscore.txt")
	if err != nil {
		rl.TraceLog(rl.LogError, err.Error())
	}

	file, err := os.Open(fileName)
	if err != nil {
		rl.TraceLog(rl.LogError, "Could not open high score file")
		return

	}

	defer file.Close()

	_, err = fmt.Fscanf(file, "%d", &g.highScore)
	if err != nil {
		rl.TraceLog(rl.LogError, "Could not read high score value from file")
		return
	}
}
