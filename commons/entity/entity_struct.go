package entity

import "github.com/f7ed0/go-multiplayer-game/commons/objects"

type Entity struct {
	Position            objects.Point
	orientation         float64
	killable            bool
	breakable           bool
	canPassThroughWalls bool
	canPassThroughHoles bool
}

func NewEntity(killable, breakable, cptw, cpth bool) Entity {
	return Entity{
		killable:            killable,
		breakable:           breakable,
		canPassThroughWalls: cptw,
		canPassThroughHoles: cpth,
		Position:            objects.Point{X: 1, Y: 1},
	}
}
