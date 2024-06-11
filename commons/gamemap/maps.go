package gamemap

import hitbox "github.com/f7ed0/go-multiplayer-game/commons/Hitbox"

type HeadLessMap struct {
	Path  string
	Walls []hitbox.Hitbox `json:"walls"`
	Holes []hitbox.Hitbox `json:"holes"`
}
