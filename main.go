package main

import (
	"fmt"
	"image/color"
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
	gravity            = -10.0
	groundScale        = 50.0
)

type Game struct {
	world    *box2d.B2World
	player   *Player
	player2  *Player
	entities []Entity
	*Map
}

func (g *Game) makeEntity(name string, bodyDef *box2d.B2BodyDef, shape *box2d.B2PolygonShape, image *ebiten.Image) (bod *box2d.B2Body) {

	bod = g.world.CreateBody(bodyDef)
	bod.CreateFixture(shape, 0.0)

	entity := Entity{
		name:    "Ground",
		bodyDef: bodyDef,
		body:    bod,
		sprite:  *image,
	}

	g.entities = append(g.entities, entity)

	return bod
}

func NewGame() *Game {
	game := Game{}
	world, entities := CreateWorld()
	game.world = &world

	game.entities = append(game.entities, entities...)

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
	shape.SetAsBox(12, 12) // A box with width=2 and height=2
	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = &shape
	fixtureDef.Density = 1.0
	fixtureDef.Friction = 0.3
	playerBody.CreateFixtureFromDef(&fixtureDef) // Create player

	game.player = player
	game.player2 = player2

	return &game
	// return &Game{
	// 	world:   &world,
	// 	player:  player,
	// 	player2: player2,
	// }
}

func (g *Game) Update() error {
	if g.player == nil {
		return nil
	}
	force := HandlePlayerInput()

	g.player.body.SetLinearVelocity(force)
	g.player2.body.SetLinearVelocity(force)

	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.player.tryJump()
	}

	g.world.Step(timeStep, velocityIterations, positionIterations)

	return nil
}

func DrawGround(screen *ebiten.Image) {
	// Ground

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Invert()

	// Draw the ground (simple rectangle)
	groundImage := ebiten.NewImage(int(screenWidth), int(-groundScale))
	groundImage.Fill(color.RGBA{255, 100, 100, 255})
	screen.DrawImage(groundImage, op)

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

	for _, element := range g.entities {

		spriteOp := ebiten.DrawImageOptions{}

		pos = element.body.GetPosition()
		spriteOp.GeoM.Translate(pos.X, pos.Y)
		screen.DrawImage(&element.sprite, &spriteOp)

	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Fixed Box2D and Ebiten Example")

	// LoadTilesetImage()
	CreateWorld()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
