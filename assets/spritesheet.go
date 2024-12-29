package assets

import "image"

type SpriteSheet struct {
	WidthInTiles  int
	HeightInTiles int
	Tilesize      int
}

func (sprite *SpriteSheet) Rect(index int) image.Rectangle {
	x := (index % sprite.WidthInTiles) * sprite.Tilesize
	y := (index % sprite.HeightInTiles) * sprite.Tilesize
	return image.Rect(
		x, y, x+sprite.Tilesize, y+sprite.Tilesize,
	)
}

func NewSpriteSheet(width, height, t int) *SpriteSheet {
	return &SpriteSheet{
		width, height, t,
	}
}
