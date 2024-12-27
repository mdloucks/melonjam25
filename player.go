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

const (
	unit    = 16
	groundY = 380
)

func NewPlayer(spritePath string, x float64, y float64, name string) (*Player, error) {

	img, _, err := ebitenutil.NewImageFromFile(spritePath)

	if err != nil {
		fmt.Printf("Could not create new player %s", err)
		defaultImg := ebiten.NewImage(96, 96)
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
			name:    name,
			bodyDef: &bodyDef,
			body:    nil,
		},
	}, nil
}

func HandlePlayerInput() box2d.B2Vec2 {
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

const (
	jumpHeight = 50
)

func (p *Player) tryJump() {
	jumpForce := box2d.MakeB2Vec2(0, jumpHeight)
	p.Entity.body.ApplyLinearImpulseToCenter(jumpForce, true)

}
