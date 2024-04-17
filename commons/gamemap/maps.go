package gamemap

import "github.com/f7ed0/go-multiplayer-game/commons/objects"

type GameMap struct {
	Walls []objects.Polygon `json:"walls"`
	Holes []objects.Polygon `json:"holes"`
}
