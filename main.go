package main

import (
	"fmt"
	"log"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth        = 640
	screenHeight       = 480
	timeStep           = 1.0 / 60.0
	velocityIterations = 6
	positionIterations = 2
)

type Game struct {
	world  *box2d.B2World
	player *Player
}

func NewGame() *Game {
	// Create the Box2D world with gravity
	gravity := box2d.MakeB2Vec2(0.0, 0.0)
	world := box2d.MakeB2World(gravity)

	// Create the ground body
	groundDef := box2d.MakeB2BodyDef()
	ground := world.CreateBody(&groundDef)

	groundShape := box2d.MakeB2EdgeShape()
	groundShape.Set(box2d.MakeB2Vec2(-20.0, 0.0), box2d.MakeB2Vec2(20.0, 0.0))
	ground.CreateFixture(&groundShape, 0.0)

	// Create player
	player, err := newPlayer("assets/img/player.png", 100, 100)

	if err != nil {
		fmt.Println("Could not create player!")
		panic("Could not create player!")
	}

	// Create the player body
	playerBody := world.CreateBody(player.bodyDef)
	player.body = playerBody

	// Attach a shape to the player body
	shape := box2d.MakeB2PolygonShape()
	shape.SetAsBox(1.0, 1.0) // A box with width=2 and height=2
	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = &shape
	fixtureDef.Density = 1.0
	fixtureDef.Friction = 0.3
	playerBody.CreateFixtureFromDef(&fixtureDef)

	return &Game{
		world:  &world,
		player: player,
	}
}

func (g *Game) Update() error {
	force := handlePlayerInput(g.player)

	g.player.body.SetLinearVelocity(force)

	g.world.Step(timeStep, velocityIterations, positionIterations)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	op := ebiten.DrawImageOptions{}

	pos := g.player.body.GetPosition()
	op.GeoM.Translate(pos.X, pos.Y)
	screen.DrawImage(&g.player.sprite, &op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Fixed Box2D and Ebiten Example")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
