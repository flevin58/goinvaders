package main

//go:generate embed -verbose -exclude_dir src -include ttf,png,xml,ogg -byte all internal/assets

import (
	"goinvaders/internal/game"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	offset       = 50
	windowWidth  = 750
	windowHeight = 700
	windowTitle  = "Golang Space Invaders"
)

func main() {

	rl.InitWindow(windowWidth+offset, windowHeight+2*offset, windowTitle)
	defer rl.CloseWindow()

	rl.InitAudioDevice()
	if !rl.IsAudioDeviceReady() {
		rl.TraceLog(rl.LogError, "Audio device not ready")
	}
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)
	rl.SetTraceLogLevel(rl.LogInfo)

	game := game.New()

	for !game.ShouldQuit() {
		rl.BeginDrawing()
		game.HandleInput()
		game.Update()
		game.Draw()
		rl.EndDrawing()
	}
}
