package main

import (
	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	name    string
	bodyDef *box2d.B2BodyDef
	body    *box2d.B2Body
	sprite  ebiten.Image
}

type Character struct {
	hp, atk, stam int
}
