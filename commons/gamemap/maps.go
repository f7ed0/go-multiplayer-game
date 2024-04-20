package gamemap

import (
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
)

type HeadLessMap struct {
	Path  string
	Walls []objects.Polygon `json:"walls"`
	Holes []objects.Polygon `json:"holes"`
}
