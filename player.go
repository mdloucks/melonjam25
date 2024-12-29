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
	isGrounded    bool
	isActive      bool
	fixture       box2d.B2FixtureDef
	hasDoubleJump bool
	hitPoints     int
}

const (
	jumpHeight   = -50000
	playerWidth  = 16
	playerHeight = 16
	maxSpeed     = 20.0
	maxJump      = 100.0
	moveForce    = 50.0
	hitPoints    = 10
	lowestPoint  = 480
)

func NewPlayer(spritePath string, x float64, y float64, name string, active bool) (*Player, error) {

	img, _, err := ebitenutil.NewImageFromFile(spritePath)

	if err != nil {
		fmt.Printf("Could not create new player %s", err)
		return DefaultPlayer(), nil
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
		false,
		hitPoints,
	}, nil
}

func (p *Player) Die(reason string) {
	// Return to Menu
	fmt.Print(p.name, "Has Died! ", reason, "\n")
	p.hitPoints = hitPoints
}
func (p *Player) CalculateDamage(damage int) {
	fmt.Print(p.name, " lost ", damage, "hp\n")
	if damage > 0 {
		p.hitPoints = p.hitPoints - damage
	}
	p.HealthCheck()

}
func (p *Player) HealthCheck() {
	if p.hitPoints <= 0 {
		p.Die("lost too much hp, git gud")
	}
}
func (p *Player) HeightCheck() {
	// Y axis is inverted, so down is positive Y
	if p.body.GetPosition().Y >= lowestPoint {
		p.Die("you fell off the map, sucks to suck")
	}
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
		p.jump()
		p.isGrounded = false
		p.hasDoubleJump = true
	} else if p.hasDoubleJump {
		p.jump()
		p.hasDoubleJump = false
	} else { // restore jumpability
		p.isGrounded = true
	}

}
func (p *Player) jump() {
	jumpForce := box2d.MakeB2Vec2(0, jumpHeight*pixlesPerMeter)
	p.Entity.body.ApplyForceToCenter(jumpForce, true)

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
	currentVelocity := p.body.GetLinearVelocity()
	l, r := ebiten.IsKeyPressed(ebiten.KeyA), ebiten.IsKeyPressed(ebiten.KeyD)
	force := box2d.MakeB2Vec2(0, 0)

	if l && !r {
		force.X = -moveForce
	} else if !l && r {
		force.X = moveForce
	} else if l && r {
		p.body.SetLinearVelocity(box2d.MakeB2Vec2(0, currentVelocity.Y))
	}
	// Max Speed
	if math.Abs(currentVelocity.X) > maxSpeed {
		currentVelocity.X = math.Copysign(maxSpeed, currentVelocity.X)
		p.body.SetLinearVelocity(box2d.MakeB2Vec2(currentVelocity.X, currentVelocity.Y))
	}
	// Max Jump, !take Abs val because fast falling is cool
	if currentVelocity.Y > maxJump {
		currentVelocity.Y = math.Copysign(maxSpeed, currentVelocity.Y)
		p.body.SetLinearVelocity(box2d.MakeB2Vec2(currentVelocity.X, currentVelocity.Y))
	}
	p.body.ApplyLinearImpulseToCenter(force, true)

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		p.tryJump()
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.body.ApplyLinearImpulseToCenter(box2d.MakeB2Vec2(0, gravity*5), true)
	}
}

func DefaultPlayer() *Player {
	defaultImg := ebiten.NewImage(playerWidth, playerHeight)
	defaultImg.Fill(color.RGBA{G: 255, A: 255})

	return &Player{
		&Entity{"", &box2d.B2BodyDef{}, &box2d.B2Body{}, *defaultImg},
		false,
		false,
		*PlayerFixture(),
		true,
		hitPoints,
	}

}
