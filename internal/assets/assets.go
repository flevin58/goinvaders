package assets

import (
	"goinvaders/internal/assets/fonts"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	Grey   = color.RGBA{R: 29, G: 29, B: 27, A: 255}
	Yellow = color.RGBA{R: 243, G: 216, B: 63, A: 255}
)

// Function to load an embedded Font
// For the purpose of this game two assumptions are made:
// - the font type is TTF
// - the font is in the assets/fonts folder of the embedded FS
func LoadFont(fontData []byte) rl.Font {
	return rl.LoadFontFromMemory(".ttf", fontData, 64, nil)
}

// Function to load an embedded Image
// For the purpose of this game two assumptions are made:
// - the image type is PNG
// - the image is in the assets/images folder of the embedded FS
func LoadImage(imageData []byte) *rl.Image {
	return rl.LoadImageFromMemory(".png", imageData, int32(len(imageData)))
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
func LoadMusic(musicData []byte) rl.Music {
	return rl.LoadMusicStreamFromMemory(".ogg", musicData, int32(len(musicData)))
}

// Function to load an embedded Sound file (sFX for laser, explosion etc.)
// For the purpose of this game two assumptions are made:
// - the sound type is OGG
// - the sound file is in the assets/sounds folder of the embedded FS
func LoadSound(soundData []byte) rl.Sound {
	wave := rl.LoadWaveFromMemory(".ogg", soundData, int32(len(soundData)))
	return rl.LoadSoundFromWave(wave)
}

func Font() rl.Font {
	return LoadFont(fonts.Monogram_ttf)
}
