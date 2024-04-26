package handleplayer

import (
	"math"
	"sync"

	"github.com/f7ed0/go-multiplayer-game/client/handleplayer/camera"
	"github.com/f7ed0/go-multiplayer-game/commons/entity/player"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/veandco/go-sdl2/sdl"
)

type HandledPlayer struct {
	Camera           camera.Camera
	PlayerController Controller
	DispPlayer
	sync.RWMutex
}

type Controller struct {
	KBActionMap map[sdl.Keycode]player.Action
}

func (hp *HandledPlayer) HandleKeyboard(sym sdl.Keycode, pressed bool) {
	action, ok := hp.PlayerController.KBActionMap[sym]
	if ok {
		if pressed {
			hp.ActionBuffer.Set(action)
		} else {
			hp.ActionBuffer.Unset(action)
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
		Camera: camera.Camera{
			Position: objects.Point{
				X: 0,
				Y: 0,
				Z: 5,
			},
		},
		DispPlayer: FromPlayerCore(player.NewPlayer()),
	}
}

func (hp *HandledPlayer) ApplyEvent(delta float32) {
	dist := 100 / hp.Camera.Position.Z
	hp.PlayerCore.ApplyEvent(delta)
	diff := objects.Diff(hp.Camera.Position, hp.Position)
	//lg.Debug.Println(diff.X, objects.Sign(diff.X))
	if math.Abs(float64(diff.X)) > float64(dist) {
		hp.Camera.Position.X = hp.Position.X - dist*objects.Sign(diff.X)
	}
	if math.Abs(float64(diff.Y)) > float64(dist) {
		hp.Camera.Position.Y = hp.Position.Y - dist*objects.Sign(diff.Y)
	}
}
