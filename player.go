package main

import (
	"fmt"
	"image/color"
	//"math"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Player struct {
	*Entity
}

const (
	unit       = 16
	groundY    = 380
	jumpHeight = -50000
)

func NewPlayer(spritePath string, x float64, y float64, name string) (*Player, error) {

	img, _, err := ebitenutil.NewImageFromFile(spritePath)

	if err != nil {
		fmt.Printf("Could not create new player %s", err)
		defaultImg := ebiten.NewImage(192, 192)
		defaultImg.Fill(color.RGBA{G: 255, A: 255})
		return &Player{&Entity{"", &box2d.B2BodyDef{}, &box2d.B2Body{}, *defaultImg}}, nil
	}

	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	bodyDef.Position.Set(x, y) // Initial position
	bodyDef.LinearDamping = 0.8

	return &Player{
		&Entity{
			name:    name,
			bodyDef: &bodyDef,
			body:    nil,
			sprite:  *img,
		},
	}, nil
}

func HandlePlayerInput() box2d.B2Vec2 {
	var force box2d.B2Vec2

	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		force.Y = 100.0 * pixlesPerMeter
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		force.Y = -100.0 * pixlesPerMeter
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		force.X = 100.0 * pixlesPerMeter
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		force.X = -100.0 * pixlesPerMeter
	}
	return force
}

func (p *Player) tryJump() {
	//velocity := p.body.GetLinearVelocity()
	//if math.Abs(velocity.Y) < 0.01 { // ground check

	//}
	jumpForce := box2d.MakeB2Vec2(0, jumpHeight*pixlesPerMeter)
	p.Entity.body.ApplyForceToCenter(jumpForce, true)

}
