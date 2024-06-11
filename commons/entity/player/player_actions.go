package player

import "math"

type Action uint64

const (
	MOVE_UP    Action = 1 << 0
	MOVE_DOWN  Action = 1 << 1
	MOVE_LEFT  Action = 1 << 2
	MOVE_RIGHT Action = 1 << 3
	SPELL_1    Action = 1 << 4
	SPELL_2    Action = 1 << 5
	SPELL_3    Action = 1 << 6
)

type ActionBuffer Action

func (ab ActionBuffer) Get(action Action) bool {
	return ab&ActionBuffer(action) > 0
}

func (ab *ActionBuffer) Set(action Action) {
	*ab |= ActionBuffer(action)
}

func (ab *ActionBuffer) Unset(action Action) {
	*ab &= ActionBuffer(math.MaxUint64 - action)
}
