package handleplayer

import (
	"math"
	"sync"

	"github.com/f7ed0/go-multiplayer-game/commons/entity/player"
	"github.com/veandco/go-sdl2/sdl"
)

type HandledPlayer struct {
	PlayerController Controller
	player.PlayerCore
	sync.RWMutex
}

type Controller struct {
	KBActionMap map[sdl.Keycode]player.Action
}

func (hp *HandledPlayer) HandleKeyboard(sym sdl.Keycode, pressed bool) {
	action, ok := hp.PlayerController.KBActionMap[sym]
	if ok {
		if pressed {
			hp.ActionBuffer |= action
		} else {
			hp.ActionBuffer &= math.MaxUint64 - action
		}
	}
}

func NewHandledPlayer() HandledPlayer {
	return HandledPlayer{
		PlayerController: Controller{
			KBActionMap: map[sdl.Keycode]player.Action{
				sdl.K_z: player.MOVE_UP,
				sdl.K_q: player.MOVE_LEFT,
				sdl.K_d: player.MOVE_RIGHT,
				sdl.K_s: player.MOVE_DOWN,
			},
		},
		PlayerCore: player.NewPlayer(),
	}
}
