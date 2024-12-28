package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func RenderMap(xOffset int, yOffset int, screen *ebiten.Image, g *Game, op *ebiten.DrawImageOptions) {

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

			op.GeoM.Translate(float64(x+xOffset), float64(y+yOffset))

			screen.DrawImage(
				g.tilemapImage.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image),
				op,
			)

			op.GeoM.Reset()
		}
	}
}

// Render an entity fit to a given size
func RenderSizedEntity(desiredW int, desiredH int, screen *ebiten.Image, entity *Entity, op *ebiten.DrawImageOptions) {
	w, h := entity.sprite.Bounds().Dx(), entity.sprite.Bounds().Dy()
	scaleX := float64(desiredW) / float64(w)
	scaleY := float64(desiredH) / float64(h)

	// Apply scaling
	op.GeoM.Scale(scaleX, scaleY)
	RenderEntity(screen, entity, op)
}

func RenderEntity(screen *ebiten.Image, entity *Entity, op *ebiten.DrawImageOptions) {

	pos := entity.body.GetPosition()
	op.GeoM.Translate(pos.X, pos.Y)
	screen.DrawImage(&entity.sprite, op)
}
