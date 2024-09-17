package assets

import (
	"embed"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

//go:embed fonts images sounds
var assetFS embed.FS

// Helper function to load any embedded file and return its data as []byte
// On error a raylib log is written and the program exits
func getEmbeddedData(filename string) []byte {
	data, err := assetFS.ReadFile(filename)
	if err != nil {
		rl.TraceLog(rl.LogError, "Could not load %s: %s", filename, err.Error())
		os.Exit(1)
	}
	return data
}

// Function to load an embedded Font
// For the purpose of this game two assumptions are made:
// - the font type is TTF
// - the font is in the assets/fonts folder of the embedded FS
func LoadFont(filename string) rl.Font {
	filename = path.Join("fonts", filename)
	data := getEmbeddedData(filename)
	return rl.LoadFontFromMemory(".ttf", data, 64, nil)
}

// Function to load an embedded Image
// For the purpose of this game two assumptions are made:
// - the image type is PNG
// - the image is in the assets/images folder of the embedded FS
func LoadImage(filename string) *rl.Image {
	filename = path.Join("images", filename)
	data := getEmbeddedData(filename)
	return rl.LoadImageFromMemory(".png", data, int32(len(data)))
}

// Function to load an image from a TextureAtlas made with TexturePacker
// It assumes that an Atlas structure is already initialized as ShipAtlas
// For more details on the Atlas refer to the atlas.go file
func LoadTexture(filename string) rl.Texture2D {
	image := rl.ImageFromImage(*ShipAtlas.Image, ShipAtlas.Sprites[filename])
	return rl.LoadTextureFromImage(&image)
}

// Function to load an embedded Music file
// For the purpose of this game two assumptions are made:
// - the music type is OGG
// - the music file is in the assets/sounds folder of the embedded FS
func LoadMusic(filename string) rl.Music {
	filename = path.Join("sounds", filename)
	data := getEmbeddedData(filename)
	return rl.LoadMusicStreamFromMemory(".ogg", data, int32(len(data)))
}

// Function to load an embedded Sound file (sFX for laser, explosion etc.)
// For the purpose of this game two assumptions are made:
// - the sound type is OGG
// - the sound file is in the assets/sounds folder of the embedded FS
func LoadSound(filename string) rl.Sound {
	filename = path.Join("sounds", filename)
	data := getEmbeddedData(filename)
	wave := rl.LoadWaveFromMemory(".ogg", data, int32(len(data)))
	return rl.LoadSoundFromWave(wave)
}
