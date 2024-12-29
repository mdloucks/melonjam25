package main

import (
	"log"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
)

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

func (g *Game) Update() error {
	switch g.state {
	case StateMenu, StateDeath:
		HandleMenu(g)
	case StatePlaying:
		HandleGameplay(g)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case StateMenu:
		DrawMenu(screen, g, "BOX2D SUCKS", "Start Game!")
	case StatePlaying:
		DrawGame(screen, g)
	case StateDeath:
		DrawMenu(screen, g, "YOU DIED!", "Try Again!")
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
