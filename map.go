package main

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	tileSize = 16
)

var (
	tilesImage *ebiten.Image
	layers     [][]int
)

type Map struct {
	layers [][]int
}

func LoadTilesetImage() {
	var err error
	tilesImage, _, err = ebitenutil.NewImageFromFile("assets/Sprite-0005.png")
	fp, err := os.Open("assets/Map.json")
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}

	// jsonData, err := io.ReadAll(fp)

	if err != nil {
		log.Fatal(err)
	}

	// layers, err = parseSpriteSheet(jsonData)
}

func (g *Game) DrawMap(screen *ebiten.Image) {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / tileSize

	const xCount = screenWidth / tileSize
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
