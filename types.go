package main

import (
	"pod/melonjam/assets"

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

type GameState int

type Game struct {
	world             *box2d.B2World
	player            *Player
	playerSpriteSheet *assets.SpriteSheet
	player2           *Player
	entities          []Entity
	tilemapJson       TilemapJSON
	tilemapImage      *ebiten.Image
	cam               Camera
	state             GameState
}

type Camera struct {
	zoom float64
	x    int
	y    int
}

type Player struct {
	*Entity
	isGrounded    bool
	isActive      bool
	fixture       box2d.B2FixtureDef
	hasDoubleJump bool
	hitPoints     int
}
