package player

import (
	"github.com/f7ed0/go-multiplayer-game/commons/entity"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
)

type Player interface {
	Move(mvtVector objects.Vector)
}

type PlayerCore struct {
	entity.Entity
}

func NewPlayer() PlayerCore {
	return PlayerCore{
		entity.NewEntity(true, false),
	}
}

func (p *PlayerCore) Move(mvtVector objects.Vector) {
	p.Position.Add(mvtVector)
}
