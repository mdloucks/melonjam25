package main

// Game Constants
const (
	screenWidth                  = 1080
	screenHeight                 = 720
	timeStep                     = 1.0 / 60.0
	velocityIterations           = 6
	positionIterations           = 2
	gravity                      = 80.0
	pixlesPerMeter               = 50
	moveSpeed                    = 100 * pixlesPerMeter
	StateMenu          GameState = iota
	StatePlaying
	StateDeath
)

// Player Constants
const (
	jumpHeight   = -50000
	playerWidth  = 64
	playerHeight = 64
	maxSpeed     = 20.0
	maxJump      = 100.0 * pixlesPerMeter
	moveForce    = 50.0 * pixlesPerMeter
	hitPoints    = 10
	lowestPoint  = 480
)
