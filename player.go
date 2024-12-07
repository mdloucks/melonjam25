package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	sprite ebiten.Image
	*Entity
}

func newPlayer(spritePath string) (*Player, error) {

	img, _, err := ebitenutil.NewImageFromFile(spritePath)

	if err != nil {
		fmt.Println("there was an error")
		fmt.Printf("%s err ", err)
		defaultImg := ebiten.NewImage(32, 32)
		defaultImg.Fill(color.White)
		return &Player{*defaultImg, &Entity{0, 0}}, nil
	}

	return &Player{*img, &Entity{0, 0}}, nil

}

func handlePlayerInput(player *Player) {
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		player.y += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		player.y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		player.x += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		player.x -= 1
	}
}
