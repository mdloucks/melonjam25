package main

import (
	"fmt"
	"image/color"
	"log"
	"pod/melonjam/animations"
	"pod/melonjam/assets"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

func HandleGameplay(g *Game) error {
	g.cam.x = int(g.player.body.GetPosition().X - (screenWidth / 4))
	g.playerIdleAnimation.Update()
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
		if element.HealthCheck() || element.HeightCheck() {
			g.state = StateDeath
		}
	}

	g.world.Step(timeStep, velocityIterations, positionIterations)

	return nil

}

// Used for Death too
func HandleMenu(g *Game) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x > 300 && y < 500 && y > 250 && y < 300 {
			g.state = StatePlaying
		}

	}
	return nil
}

func DrawGame(screen *ebiten.Image, g *Game) {
	screen.Fill(color.RGBA{60, 100, 150, 255})

	op := ebiten.DrawImageOptions{}

	RenderMap(0, 0, screen, g, &op, &g.cam)
	RenderMap(0, 200, screen, g, &op, &g.cam)

	op.GeoM.Reset()
	RenderSizedEntity(playerWidth, playerHeight, screen, g.player.Entity, &op, &g.cam, g.playerIdleAnimation.Frame())

	op.GeoM.Reset()
	RenderSizedEntity(playerWidth, playerHeight, screen, g.player2.Entity, &op, &g.cam, g.playerIdleAnimation.Frame())

	op.GeoM.Reset()

	for _, element := range g.entities {
		RenderEntity(screen, &element, &op, &g.cam, g.playerIdleAnimation.Frame())
		op.GeoM.Reset()
	}

}
func DrawMenu(screen *ebiten.Image, g *Game, topText string, bottomText string) {

	screen.Fill(color.Black)

	face := basicfont.Face7x13
	text.Draw(screen, topText, face, 340, 150, color.White)

	vector.DrawFilledRect(screen, 300, 250, 200, 50, color.RGBA{R: 0, G: 255, B: 0, A: 255}, false) // Green button
	text.Draw(screen, bottomText, face, 335, 280, color.Black)
}

func NewGame() *Game {
	playerSpriteSheet := assets.NewSpriteSheet(8, 1, 64)
	playerIdleAnimation := animations.NewAnimation(0, 7, 1, 20.0)
	game := Game{}
	// init in the menu
	game.state = StateMenu
	game.playerSpriteSheet = playerSpriteSheet
	game.playerIdleAnimation = playerIdleAnimation

	world, entities := CreateWorld()
	game.world = &world

	game.entities = append(game.entities, entities...)

	player, err := NewPlayer("assets/img/GreenManIdle.png", 100, 300, "dark", true, playerSpriteSheet)

	player2, err2 := NewPlayer("assets/img/GreenManIdle.png", 100, 150, "light", false, playerSpriteSheet)

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
		y:    80,
		zoom: 2.0,
	}

	game.cam = cam

	return &game
}
