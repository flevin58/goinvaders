package assets

import (
	"encoding/xml"
	"fmt"
	"os"
	"path"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// The following two structures map the atlas xml file produced by TexturePacker
// For detail look at the ships.xml file, it is quite straightforward to understand
type sprite struct {
	Name string `xml:"n,attr"`
	X    uint   `xml:"x,attr"`
	Y    uint   `xml:"y,attr"`
	W    uint   `xml:"w,attr"`
	H    uint   `xml:"h,attr"`
}

type textureAtlas struct {
	Sprites []sprite `xml:"sprite"`
}

// This is the published Atlas structure that is used to get sprite images
// Image is the actual SpriteSheet png file with all the sprites
// Sprites is a map that returns an rl.Rectangle for each image name (the original filename)
// It is used to get the Sub-Textures of each sprite from the full texture
type Atlas struct {
	Image   *rl.Image
	Sprites map[string]rl.Rectangle
}

var ShipAtlas = NewAtlas("ships")

func NewAtlas(atlasName string) *Atlas {

	xmlFile := path.Join("images", atlasName+".xml")
	xmlData := getEmbeddedData(xmlFile)
	Ta := textureAtlas{}
	err := xml.Unmarshal(xmlData, &Ta)
	if err != nil {
		rl.TraceLog(rl.LogError, "Could not unmarshal %s: %s", xmlFile, err.Error())
		os.Exit(1)
	}

	pngFile := atlasName + ".png"
	atlas := &Atlas{
		Image:   LoadImage(pngFile),
		Sprites: make(map[string]rl.Rectangle),
	}
	for _, s := range Ta.Sprites {
		atlas.Sprites[s.Name] = rl.Rectangle{
			X:      float32(s.X),
			Y:      float32(s.Y),
			Width:  float32(s.W),
			Height: float32(s.H),
		}
	}

	return atlas
}

// The following functions get the respective image textures
// They actually hide the file system structure so to create a layer of
// abstraction, giving freedom to move images around without breakin the code.

func GetAlienImage(alienType int32) rl.Texture2D {
	name := fmt.Sprintf("alien_%d.png", alienType)
	return LoadTexture(name)
}

func GetSpaceshipImage() rl.Texture2D {
	return LoadTexture("spaceship.png")
}
func GetMysteryImage() rl.Texture2D {
	return LoadTexture("mystery.png")
}
