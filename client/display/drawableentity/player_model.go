package drawableentity

import (
	"github.com/f7ed0/go-multiplayer-game/client/animatedsprite"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
)

type EntityModel struct {
	States   map[string]animatedsprite.AnimatedSptite
	ViewPort objects.Vector
}
