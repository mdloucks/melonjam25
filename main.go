package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth        = 640
	screenHeight       = 480
	timeStep           = 1.0 / 60.0
	velocityIterations = 6
	positionIterations = 2
	gravity            = 80.0
	pixlesPerMeter     = 50
	moveSpeed          = 100 * pixlesPerMeter
)

type Game struct {
	world        *box2d.B2World
	player       *Player
	player2      *Player
	entities     []Entity
	tilemapJson  TilemapJSON
	tilemapImage *ebiten.Image
}

func (g *Game) makeEntity(name string, bodyDef *box2d.B2BodyDef, fixtureDef *box2d.B2FixtureDef, image *ebiten.Image) (bod *box2d.B2Body) {

	bod = g.world.CreateBody(bodyDef)
	bod.CreateFixtureFromDef(fixtureDef)

	entity := Entity{
		name:    name,
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

	player, err := NewPlayer("assets/img/player.png", 100, 350, "dark")

	player2, err2 := NewPlayer("assets/img/player.png", 100, 450, "light")

	if err != nil || err2 != nil {
		fmt.Println("Could not create player!")
		panic("Could not create player!")
	}

	// Create the player body
	player.bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	playerBody := world.CreateBody(player.bodyDef)
	player.body = playerBody

	player2Body := world.CreateBody(player2.bodyDef)
	player2.body = player2Body

	// Attach a shape to the player body
	shape := box2d.MakeB2PolygonShape()
	shape.SetAsBox(0.5, 0.5) // A box with width=2 and height=2

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
	fixtureDef.Density = 1.0
	fixtureDef.Friction = 0.3
	fixtureDef.Restitution = 0.0
	playerBody.CreateFixtureFromDef(&fixtureDef) // Create player

	game.player = player
	game.player2 = player2

	createMap(&game, &world)

	return &game
}

func (g *Game) Update() error {
	if g.player == nil {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyL) {
		force := box2d.MakeB2Vec2(moveSpeed, 0)
		g.player.body.ApplyLinearImpulseToCenter(force, true)
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		force := box2d.MakeB2Vec2(-moveSpeed, 0)
		g.player.body.ApplyLinearImpulseToCenter(force, true)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.player.tryJump()
	}

	g.world.Step(timeStep, velocityIterations, positionIterations)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{60, 100, 150, 255})

	op := ebiten.DrawImageOptions{}

	for _, layer := range g.tilemapJson.Layers {

		for index, id := range layer.Data {

			x := index % layer.Width
			y := index / layer.Width

			x *= 16
			y *= 16

			srcX := (id - 1) % 22
			srcY := (id - 1) / 22

			srcX *= 16
			srcY *= 16

			op.GeoM.Translate(float64(x), float64(y))

			screen.DrawImage(
				g.tilemapImage.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image),
				&op,
			)

			op.GeoM.Reset()
		}
	}

	op.GeoM.Reset()

	pos := g.player.body.GetPosition()
	op.GeoM.Translate(pos.X, pos.Y)
	screen.DrawImage(&g.player.sprite, &op)

	op.GeoM.Reset()
	pos2 := g.player2.body.GetPosition()
	op.GeoM.Translate(pos2.X, pos2.Y)
	screen.DrawImage(&g.player2.sprite, &op)

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
	ebiten.SetWindowTitle("BOX2D SUCKS")

	// LoadTilesetImage()
	CreateWorld()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
