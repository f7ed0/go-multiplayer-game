package player

import (
	"github.com/f7ed0/go-multiplayer-game/commons/entity"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
)

type Player interface {
	Move(mvtVector objects.Vector)
}

type Action uint64

type PlayerCore struct {
	ActionBuffer Action
	entity.Entity
}

func NewPlayer() PlayerCore {
	return PlayerCore{
		ActionBuffer: 0,
		Entity:       entity.NewEntity(true, false),
	}
}

func (p *PlayerCore) Move(mvtVector objects.Vector) {
	p.Position.Add(mvtVector)
}

const (
	MOVE_UP    Action = 1 << 0
	MOVE_DOWN  Action = 1 << 1
	MOVE_LEFT  Action = 1 << 2
	MOVE_RIGHT Action = 1 << 3
	SPELL_1    Action = 1 << 4
	SPELL_2    Action = 1 << 5
	SPELL_3    Action = 1 << 6
)

func (p *PlayerCore) ApplyEvent(delta float32) {
	if p.ActionBuffer&MOVE_UP > 0 {
		p.Position.Y -= 100 * delta
	}
	if p.ActionBuffer&MOVE_DOWN > 0 {
		p.Position.Y += 100 * delta
	}
	if p.ActionBuffer&MOVE_LEFT > 0 {
		p.Position.X -= 100 * delta
	}
	if p.ActionBuffer&MOVE_RIGHT > 0 {
		p.Position.X += 100 * delta
	}
}
