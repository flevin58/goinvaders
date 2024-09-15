package assets

import (
	"embed"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed fonts images sounds
var assetFS embed.FS

func LoadFont(filename string) rl.Font {
	filename = path.Join("internal", "assets", "fonts", filename)
	return rl.LoadFontEx(filename, 64, nil, 0)
}

func LoadTexture(filename string) rl.Texture2D {
	filename = path.Join("internal", "assets", "images", filename)
	//filename = path.Join("images", filename)
	return rl.LoadTexture(filename)
}

func LoadMusic(filename string) rl.Music {
	filename = path.Join("internal", "assets", "sounds", filename)
	return rl.LoadMusicStream(filename)
}

func LoadSound(filename string) rl.Sound {
	filename = path.Join("internal", "assets", "sounds", filename)
	return rl.LoadSound(filename)
}
