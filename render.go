package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

func RenderMap(xOffset int, yOffset int, screen *ebiten.Image, g *Game, op *ebiten.DrawImageOptions, cam *Camera) {
	for _, layer := range g.tilemapJson.Layers {
		for index, id := range layer.Data {
			if id == 0 {
				continue // Skip empty tiles
			}

			// Calculate tile position
			x := index % layer.Width
			y := index / layer.Width

			// Tile size is 16x16, scale by offset
			x *= 16
			y *= 16

			// Calculate source tile coordinates in the tileset
			srcX := (id - 1) % 22
			srcY := (id - 1) / 22

			srcX *= 16
			srcY *= 16

			// Adjust tile position based on camera and offsets
			screenX := (float64(x) - float64(cam.x) + float64(xOffset)) * cam.zoom
			screenY := (float64(y) - float64(cam.y) + float64(yOffset)) * cam.zoom

			// Apply scaling and translation
			op.GeoM.Reset() // Reset before applying new transformations
			op.GeoM.Scale(cam.zoom, cam.zoom)
			op.GeoM.Translate(screenX, screenY)

			// Draw the tile
			screen.DrawImage(
				g.tilemapImage.SubImage(image.Rect(srcX, srcY, srcX+16, srcY+16)).(*ebiten.Image),
				op,
			)
		}
	}
}

// Render an entity fit to a given size, considering the camera
func RenderSizedEntity(desiredW int, desiredH int, screen *ebiten.Image, entity *Entity, op *ebiten.DrawImageOptions, cam *Camera, frame int) {
	w, h := 64, entity.sprite.Bounds().Dy()
	scaleX := (float64(desiredW) / float64(w)) * cam.zoom
	scaleY := (float64(desiredH) / float64(h)) * cam.zoom

	// Apply scaling
	op.GeoM.Scale(scaleX, scaleY)
	RenderEntity(screen, entity, op, cam, frame)
}

// Render an entity considering the camera
func RenderEntity(screen *ebiten.Image, entity *Entity, op *ebiten.DrawImageOptions, cam *Camera, frame int) {
	pos := entity.body.GetPosition()

	// Adjust position based on the camera
	worldX := (pos.X - float64(cam.x)) * cam.zoom
	worldY := (pos.Y - float64(cam.y)) * cam.zoom

	// Apply translation
	op.GeoM.Translate(worldX, worldY)
	currentImage := entity.spriteSheet
	if currentImage == nil {
		currentImage = DefaultPlayer().spriteSheet
	}
	screen.DrawImage(entity.sprite.SubImage(currentImage.Rect(frame)).(*ebiten.Image), op)
}
