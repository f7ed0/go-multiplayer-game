package player

import (
	"time"

	"github.com/f7ed0/go-multiplayer-game/commons/entity"
)

type PlayerType uint

const (
	KNIGHT PlayerType = 1
)

type PlayerCore struct {
	Type         PlayerType
	ActionBuffer ActionBuffer
	entity.Entity
	LastTime time.Time
}

func NewPlayer() PlayerCore {
	return PlayerCore{
		Type:         KNIGHT,
		ActionBuffer: 0,
		Entity:       entity.NewEntity(true, false, false, false),
		LastTime:     time.Now(),
	}
}

func (p *PlayerCore) ApplyEvent(delta float32) {
	if p.ActionBuffer.Get(MOVE_UP) {
		p.Position.Y -= 100 * delta
	}
	if p.ActionBuffer.Get(MOVE_DOWN) {
		p.Position.Y += 100 * delta
	}
	if p.ActionBuffer.Get(MOVE_LEFT) {
		p.Position.X -= 100 * delta
	}
	if p.ActionBuffer.Get(MOVE_RIGHT) {
		p.Position.X += 100 * delta
	}

}
