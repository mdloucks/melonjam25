package main

import "github.com/ByteArena/box2d"

type Entity struct {
	name    string
	bodyDef *box2d.B2BodyDef
	body    *box2d.B2Body
}

type Character struct {
	hp, atk, stam int
}
