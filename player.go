package main

import (
	"fmt"
	"image/color"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	sprite ebiten.Image
	*Entity
}

func newPlayer(spritePath string, x float64, y float64) (*Player, error) {

	img, _, err := ebitenutil.NewImageFromFile(spritePath)

	if err != nil {
		fmt.Printf("Could not create new player %s", err)
		defaultImg := ebiten.NewImage(32, 32)
		defaultImg.Fill(color.RGBA{G: 255, A: 255})
		return &Player{*defaultImg, &Entity{"", &box2d.B2BodyDef{}, &box2d.B2Body{}}}, nil
	}

	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	bodyDef.Position.Set(x, y) // Initial position
	bodyDef.LinearDamping = 0.8

	return &Player{
		*img,
		&Entity{
			name:    "",
			bodyDef: &bodyDef,
			body:    nil,
		},
	}, nil
}

func handlePlayerInput(player *Player) box2d.B2Vec2 {
	var force box2d.B2Vec2

	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		force.Y = 100.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		force.Y = -100.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		force.X = 100.0
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		force.X = -100.0
	}
	return force
}
