package main

import (
	"fmt"
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
	cam          Camera
}

type Camera struct {
	zoom float64
	x    int
	y    int
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

	player, err := NewPlayer("assets/img/GreenMan.png", 100, 300, "dark", true)

	player2, err2 := NewPlayer("assets/img/GreenMan.png", 100, 150, "light", false)

	if err != nil || err2 != nil {
		fmt.Println("Could not create player!")
		panic("Could not create player!")
	}

	// Create the player body
	player.body = world.CreateBody(player.bodyDef)
	player2.body = world.CreateBody(player2.bodyDef)

	// player fixture

	player.body.CreateFixtureFromDef(PlayerFixture()) // Create player
	player2.body.CreateFixtureFromDef(PlayerFixture())

	game.player = player
	game.player2 = player2

	createMap(0, 0, &game, &world)
	createMap(0, 200, &game, &world)

	cam := Camera{
		x:    int(player.body.GetPosition().X),
		y:    0,
		zoom: 1,
	}

	game.cam = cam

	return &game
}

func (g *Game) Update() error {
	g.cam.x = int(g.player.body.GetPosition().X - (screenWidth / 2))
	// g.cam.y = int(g.player.body.GetPosition().Y - screenHeight/2)

	var players = []*Player{g.player, g.player2}

	if g.player == nil || g.player2 == nil {
		log.Fatal(nil)
	}
	if g.player.isActive {
		HandleInput(g.player)
	} else {
		HandleInput(g.player2)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.player.swap()
		g.player2.swap()
	}

	// Damage check
	// Simulate Damage
	// g.player.CalculateDamage(1)

	for _, element := range players {
		element.HealthCheck()
		element.HeightCheck()
	}

	g.world.Step(timeStep, velocityIterations, positionIterations)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	screen.Fill(color.RGBA{60, 100, 150, 255})

	op := ebiten.DrawImageOptions{}

	RenderMap(0, 0, screen, g, &op, &g.cam)
	RenderMap(0, 200, screen, g, &op, &g.cam)

	op.GeoM.Reset()
	RenderSizedEntity(playerWidth, playerHeight, screen, g.player.Entity, &op, &g.cam)

	op.GeoM.Reset()
	RenderSizedEntity(playerWidth, playerHeight, screen, g.player2.Entity, &op, &g.cam)

	op.GeoM.Reset()

	for _, element := range g.entities {
		RenderEntity(screen, &element, &op, &g.cam)
		op.GeoM.Reset()
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
