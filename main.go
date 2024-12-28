package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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

func (g *Game) makeEntity(name string, bodyDef *box2d.B2BodyDef, shape *box2d.B2PolygonShape, image *ebiten.Image) (bod *box2d.B2Body) {

	bod = g.world.CreateBody(bodyDef)
	bod.CreateFixture(shape, 0.0)

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

	player, err := NewPlayer("assets/img/GreenMan.png", 100, 300, "dark", true)

	player2, err2 := NewPlayer("assets/img/GreenMan.png", 100, 150, "light", false)

	if err != nil || err2 != nil {
		fmt.Println("Could not create player!")
		panic("Could not create player!")
	}

	// Create the player body
	playerBody := world.CreateBody(player.bodyDef)
	player.body = playerBody

	player2Body := world.CreateBody(player2.bodyDef)
	player2.body = player2Body

	// player fixture

	playerBody.CreateFixtureFromDef(PlayerFixture()) // Create player
	player2Body.CreateFixtureFromDef(PlayerFixture())

	tilemapJson, err := NewTilemapJSON("assets/level1tilemap.tmj")
	game.tilemapJson = *tilemapJson

	if err != nil {
		log.Fatal("Could not load tilemap json")
	}

	tilemapImg, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}
	game.tilemapImage = tilemapImg

	game.player = player
	game.player2 = player2

	return &game
}

func (g *Game) Update() error {
	if g.player == nil || g.player2 == nil {
		log.Fatal(nil)
	}
	if g.player.isActive {
		HandleInput(g.player)
	} else {
		HandleInput(g.player2)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.player.swap()
		g.player2.swap()
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
