package assets

import (
	"embed"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed fonts images sounds
var assetFS embed.FS

func getEmbeddedData(filename string) []byte {
	data, err := assetFS.ReadFile(filename)
	if err != nil {
		rl.TraceLog(rl.LogError, "Could not load %s: %s", filename, err.Error())
		os.Exit(1)
	}
	return data
}

func LoadFont(filename string) rl.Font {
	filename = path.Join("fonts", filename)
	data := getEmbeddedData(filename)
	return rl.LoadFontFromMemory(".ttf", data, 64, nil)
}

func LoadTexture(filename string) rl.Texture2D {
	filename = path.Join("images", filename)
	data := getEmbeddedData(filename)
	image := rl.LoadImageFromMemory(".png", data, int32(len(data)))
	return rl.LoadTextureFromImage(image)
}

func LoadMusic(filename string) rl.Music {
	filename = path.Join("sounds", filename)
	data := getEmbeddedData(filename)
	return rl.LoadMusicStreamFromMemory(".ogg", data, int32(len(data)))
}

func LoadSound(filename string) rl.Sound {
	filename = path.Join("sounds", filename)
	data := getEmbeddedData(filename)
	wave := rl.LoadWaveFromMemory(".ogg", data, int32(len(data)))
	return rl.LoadSoundFromWave(wave)
}
