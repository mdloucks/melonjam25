package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/ByteArena/box2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func CreateWorld() (box2d.B2World, []Entity) {

	gravity := box2d.MakeB2Vec2(0.0, gravity)
	world := box2d.MakeB2World(gravity)

	var entities []Entity
	// entities = append(entities, createGround(&world, 0.0, 400.0), createGround(&world, 0.0, 200.0))

	return world, entities
}

func createGround(world *box2d.B2World, x, y float64) Entity {

	// Create the ground body
	w, h := 800.0, 30.0
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

func createMap(x int, y int, game *Game, world *box2d.B2World) {

	tilemapJson, err := NewTilemapJSON("assets/level1tilemap.tmj")
	game.tilemapJson = *tilemapJson

	if err != nil {
		log.Fatal("Could not load tilemap json")
	}

	tilemapImg, _, err := ebitenutil.NewImageFromFile("assets/images/TilesetFloor.png")
	if err != nil {
		log.Fatal(err)
	}
	game.tilemapImage = tilemapImg

	objectBodyDef := box2d.MakeB2BodyDef()
	objectBodyDef.Type = box2d.B2BodyType.B2_staticBody
	var vertices []box2d.B2Vec2

	for _, layer := range game.tilemapJson.Layers {
		for _, object := range layer.Objects {

			objectBodyDef.Position.Set(float64(object.X+x), float64(object.Y+y))

			if object.Polygon == nil {

				fmt.Printf("Creating box object... %v %v %v %v\n", object.X, object.Y, object.Width, object.Height)

				vertices = []box2d.B2Vec2{
					box2d.MakeB2Vec2(0, 0),                                          // bottom-left corner (relative to the body's position)
					box2d.MakeB2Vec2(float64(object.Width), 0),                      // bottom-right corner (relative to the body's position)
					box2d.MakeB2Vec2(float64(object.Width), float64(object.Height)), // top-right corner (relative to the body's position)
					box2d.MakeB2Vec2(0, float64(object.Height)),                     // top-left corner (relative to the body's position)
				}
			} else {

				fmt.Printf("Creating polygon object... %v \n", object.Polygon)
				vertices = []box2d.B2Vec2{}
				for _, coords := range object.Polygon {
					vertices = append(vertices, box2d.MakeB2Vec2(float64(coords.X), float64(coords.Y)))
				}
			}

			shape := box2d.MakeB2PolygonShape()
			shape.Set(vertices, len(vertices))

			fixtureDef := box2d.MakeB2FixtureDef()
			fixtureDef.Shape = &shape

			// UNCOMMENT TO DEBUG WALLS
			// var wallImage *ebiten.Image
			// if object.Polygon == nil {
			//
			// 	wallImage = ebiten.NewImage(int(object.Width), int(object.Height))
			// 	wallImage.Fill(color.CMYK{100, 200, 30, 1})
			// } else {
			//
			// 	wallImage = ebiten.NewImage(16, 16)
			// 	wallImage.Fill(color.CMYK{100, 200, 30, 1})
			// }

			wallImage := ebiten.NewImage(1, 1)
			wallImage.Fill(color.RGBA{0, 0, 0, 255})

			wallBody := game.makeEntity("Wall", &objectBodyDef, &fixtureDef, wallImage)
			fmt.Printf("Created new wall at %v %v\n", wallBody.GetPosition().X, wallBody.GetPosition().Y)
		}
	}
}
