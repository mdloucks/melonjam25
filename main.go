package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player Player
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		g.player.y += 1
	} else if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.player.y -= 1
	} else if ebiten.IsKeyPressed(ebiten.KeyL) {
		g.player.x += 1
	} else if ebiten.IsKeyPressed(ebiten.KeyH) {
		g.player.x -= 1
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.player.x, g.player.y)
	screen.DrawImage(&g.player.sprite, &op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	player, err := newPlayer("assets/img/player.png")
	game := &Game{}

	if err != nil {
		fmt.Println("could not create player!")
	}

	game.player = *player

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Melon Jam 25")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
