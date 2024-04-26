package entity

import (
	hitbox "github.com/f7ed0/go-multiplayer-game/commons/Hitbox"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
)

type Entity struct {
	Position            objects.Point
	orientation         float64
	killable            bool
	breakable           bool
	Health              float64
	canPassThroughWalls bool
	canPassThroughHoles bool
	HitBoxes            []hitbox.Hitbox
}

func NewEntity(killable, breakable, cptw, cpth bool, hits []hitbox.Hitbox) Entity {
	return Entity{
		killable:            killable,
		breakable:           breakable,
		canPassThroughWalls: cptw,
		canPassThroughHoles: cpth,
		Position:            objects.Point{X: 1, Y: 1},
		Health:              1,
		HitBoxes:            hits,
	}
}

type EntityType uint

const (
	KNIGHT EntityType = 1
)
