package main

//
// import (
// 	"fmt"
// 	"github.com/ByteArena/box2d"
// )
//
// func main() {
// 	// Define gravity (downwards)
// 	gravity := box2d.MakeB2Vec2(1.0, -10.0)
//
// 	// Create a world with the specified gravity
// 	world := box2d.MakeB2World(gravity)
//
// 	// Create the ground body (static)
// 	{
// 		bd := box2d.MakeB2BodyDef()
// 		ground := world.CreateBody(&bd)
//
// 		// Create a ground shape (horizontal line)
// 		shape := box2d.MakeB2EdgeShape()
// 		shape.Set(box2d.MakeB2Vec2(-20.0, 0.0), box2d.MakeB2Vec2(20.0, 0.0))
// 		ground.CreateFixture(&shape, 0.0)
// 	}
//
// 	// Create a dynamic box (falling object)
// 	{
// 		bd := box2d.MakeB2BodyDef()
// 		bd.Position.Set(1.0, 10.0) // Initial position above ground
// 		bd.Type = box2d.B2BodyType.B2_dynamicBody
//
// 		// Create the falling box shape
// 		body := world.CreateBody(&bd)
// 		shape := box2d.MakeB2PolygonShape()
// 		shape.SetAsBox(1.0, 1.0) // 2x2 box
// 		fd := box2d.MakeB2FixtureDef()
// 		fd.Shape = &shape
// 		fd.Density = 1.0
// 		body.CreateFixtureFromDef(&fd)
//
// 		// Run the simulation for 60 steps
// 		for {
// 			world.Step(1.0/6600.0, 9000, 3000) // Simulate one step (1/60s)
//
// 			// Get the position of the falling box and print it
// 			position := body.GetPosition()
// 			fmt.Printf("Step: Position = (%4.3f, %4.3f)\n", position.X, position.Y)
// 		}
// 	}
// }
