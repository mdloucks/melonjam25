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
	world   *box2d.B2World
	player  *Player
	player2 *Player
}

func CreateWorld() box2d.B2World {
	gravity := box2d.MakeB2Vec2(0.0, 0.0)
	world := box2d.MakeB2World(gravity)

	// Create the ground body
	groundDef := box2d.MakeB2BodyDef()
	ground := world.CreateBody(&groundDef)

	groundShape := box2d.MakeB2EdgeShape()
	groundShape.Set(box2d.MakeB2Vec2(-20.0, 0.0), box2d.MakeB2Vec2(20.0, 0.0))
	ground.CreateFixture(&groundShape, 0.0)

	return world
}

func NewGame() *Game {
	// Create the Box2D world with gravity
	world := CreateWorld()

	player, err := NewPlayer("assets/img/player.png", 100, 100, "dark")

	player2, err2 := NewPlayer("assets/img/player.png", 100, 150, "light")

	if err != nil || err2 != nil {
		fmt.Println("Could not create player!")
		panic("Could not create player!")
	}

	// Create the player body
	playerBody := world.CreateBody(player.bodyDef)
	player.body = playerBody

	player2Body := world.CreateBody(player2.bodyDef)
	player2.body = player2Body

	// Attach a shape to the player body
	shape := box2d.MakeB2PolygonShape()
	shape.SetAsBox(1.0, 1.0) // A box with width=2 and height=2
	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = &shape
	fixtureDef.Density = 1.0
	fixtureDef.Friction = 0.3
	playerBody.CreateFixtureFromDef(&fixtureDef) // Create player

	return &Game{
		world:   &world,
		player:  player,
		player2: player2,
	}
}

func (g *Game) Update() error {
	if g.player == nil {
		return nil
	}
	force := HandlePlayerInput()

	g.player.body.SetLinearVelocity(force)
	g.player2.body.SetLinearVelocity(force)

	g.world.Step(timeStep, velocityIterations, positionIterations)

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.player.tryJump()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	op := ebiten.DrawImageOptions{}

	pos := g.player.body.GetPosition()
	op.GeoM.Translate(pos.X, pos.Y)
	screen.DrawImage(&g.player.sprite, &op)

	op2 := ebiten.DrawImageOptions{}

	pos2 := g.player2.body.GetPosition()
	op2.GeoM.Translate(pos2.X, pos2.Y)
	screen.DrawImage(&g.player2.sprite, &op2)
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
