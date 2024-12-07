package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	sprite ebiten.Image
	x      float64
	y      float64
}

func newPlayer(spritePath string) (*Player, error) {

	img, _, err := ebitenutil.NewImageFromFile(spritePath)

	if err != nil {
		fmt.Println("there was an error")
		fmt.Printf("%s err ", err)
		defaultImg := ebiten.NewImage(32, 32)
		defaultImg.Fill(color.White)
		return &Player{*defaultImg, 0, 0}, nil
	}

	return &Player{*img, 0, 0}, nil

}
