package player

import (
	"time"

	"github.com/f7ed0/go-multiplayer-game/commons/entity"
)

type PlayerCore struct {
	Type         entity.EntityType
	ActionBuffer ActionBuffer
	entity.Entity
	LastTime time.Time
	Hash     string
}

func NewPlayer() PlayerCore {
	return PlayerCore{
		Type:         entity.KNIGHT,
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

func (p *PlayerCore) ClearEvent() {

}
