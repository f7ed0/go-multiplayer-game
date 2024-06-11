package player

import (
	"time"

	hitbox "github.com/f7ed0/go-multiplayer-game/commons/Hitbox"
	"github.com/f7ed0/go-multiplayer-game/commons/entity"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
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
		Entity: entity.NewEntity(true, false, false, false, []hitbox.Hitbox{
			{
				Polygon: objects.Polygon{
					Points: []objects.Point{
						{X: -8, Y: 15},
						{X: 9, Y: 15},
						{X: 9, Y: 31},
						{X: -8, Y: 31},
					},
				},
			},
		}),
		LastTime: time.Now(),
	}
}

func (p *PlayerCore) ApplyEvent(delta float32) {
	var direction objects.Vector
	if p.ActionBuffer.Get(MOVE_UP) {
		direction.Y -= 1
	}
	if p.ActionBuffer.Get(MOVE_DOWN) {
		direction.Y += 1
	}
	if p.ActionBuffer.Get(MOVE_LEFT) {
		direction.X -= 1
	}
	if p.ActionBuffer.Get(MOVE_RIGHT) {
		direction.X += 1
	}
	norm := direction.Normalized2D()
	p.Position = p.Position.Add(norm.Times(150 * delta))
}

func (p *PlayerCore) ClearEvent() {

}
