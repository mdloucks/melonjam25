package main

import (
	"image/color"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
)

func CreateWorld() (box2d.B2World, []Entity) {

	gravity := box2d.MakeB2Vec2(0.0, gravity)
	world := box2d.MakeB2World(gravity)

	var entities []Entity
	entities = append(entities, createGround(&world))

	return world, entities
}

func createGround(world *box2d.B2World) Entity {

	// Create the ground body
	x, y, w, h := 0.0, 400.0, 800.0, 30.0
	groundDef := box2d.MakeB2BodyDef()
	groundDef.Position.Set(x, y)
	groundImage := ebiten.NewImage(int(w), int(h))
	groundImage.Fill(color.CMYK{100, 200, 30, 1})
	groundShape := box2d.MakeB2PolygonShape()

	vertices := []box2d.B2Vec2{
		box2d.MakeB2Vec2(0, 0), // bottom-left corner (relative to the body's position)
		box2d.MakeB2Vec2(w, 0), // bottom-right corner (relative to the body's position)
		box2d.MakeB2Vec2(w, h), // top-right corner (relative to the body's position)
		box2d.MakeB2Vec2(0, h), // top-left corner (relative to the body's position)
	}

	groundShape.Set(vertices, len(vertices))

	groundBody := world.CreateBody(&groundDef)

	groundBody.SetType(box2d.B2BodyType.B2_staticBody)
	groundBody.CreateFixture(&groundShape, 0.0)

	entity := Entity{
		name:    "Ground",
		bodyDef: &groundDef,
		body:    groundBody,
		sprite:  *groundImage}

	return entity
}
