package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	*Entity
	isGrounded bool
	isActive   bool
	fixture    box2d.B2FixtureDef
}

const (
	jumpHeight   = -50000
	playerWidth  = 16
	playerHeight = 16
)

func NewPlayer(spritePath string, x float64, y float64, name string, active bool) (*Player, error) {

	img, _, err := ebitenutil.NewImageFromFile(spritePath)

	if err != nil {
		fmt.Printf("Could not create new player %s", err)
		defaultImg := ebiten.NewImage(playerWidth, playerHeight)
		defaultImg.Fill(color.RGBA{G: 255, A: 255})
		return &Player{&Entity{"", &box2d.B2BodyDef{}, &box2d.B2Body{}, *defaultImg}, false, false, *PlayerFixture()}, nil
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
		false,
		active,
		*PlayerFixture(),
	}, nil
}
func (p *Player) swap() {
	p.isActive = !p.isActive

	if p.isActive {
		p.body.SetGravityScale(1.0)
		p.body.CreateFixtureFromDef(&p.fixture)
	} else {
		p.body.SetGravityScale(0.0)
		p.body.DestroyFixture(p.body.GetFixtureList())
	}

}
func (p *Player) tryJump() {
	velocity := p.body.GetLinearVelocity()
	if math.Abs(velocity.Y) < 0.01 { // ground check
		jumpForce := box2d.MakeB2Vec2(0, jumpHeight*pixlesPerMeter)
		p.Entity.body.ApplyForceToCenter(jumpForce, true)
		p.isGrounded = false
	} else {
		p.isGrounded = true
	}

}

func PlayerFixture() *box2d.B2FixtureDef {
	shape := box2d.MakeB2PolygonShape()

	w, h := 16.0, 16.0
	vertices := []box2d.B2Vec2{
		box2d.MakeB2Vec2(0, 0), // bottom-left corner (relative to the body's position)
		box2d.MakeB2Vec2(w, 0), // bottom-right corner (relative to the body's position)
		box2d.MakeB2Vec2(w, h), // top-right corner (relative to the body's position)
		box2d.MakeB2Vec2(0, h), // top-left corner (relative to the body's position)
	}

	shape.Set(vertices, len(vertices))

	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = &shape
	fixtureDef.Density = 0.1
	fixtureDef.Friction = 0.3
	fixtureDef.Restitution = 0.0

	return &fixtureDef
}

func HandleInput(p *Player) {
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		force := box2d.MakeB2Vec2(moveSpeed, 0)
		p.body.ApplyLinearImpulseToCenter(force, true)
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		force := box2d.MakeB2Vec2(-moveSpeed, 0)
		p.body.ApplyLinearImpulseToCenter(force, true)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.tryJump()
	}
}
