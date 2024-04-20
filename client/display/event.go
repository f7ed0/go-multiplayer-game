package display

import (
	"math"

	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/veandco/go-sdl2/sdl"
)

func (w *Window) event(delta float32) {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.GetType() {
		case sdl.QUIT:
			w.exit = true
		case sdl.KEYDOWN:
			w.handleKeyboardEvent(e.(*sdl.KeyboardEvent), true)
		case sdl.KEYUP:
			w.handleKeyboardEvent(e.(*sdl.KeyboardEvent), false)
		}
	}
	w.Me.Lock()
	w.Me.ApplyEvent(delta)
	w.Me.Unlock()
	diff := objects.Diff(w.camera.Position, w.Me.Position)
	//lg.Debug.Println(diff.X, objects.Sign(diff.X))
	if math.Abs(float64(diff.X)) > 50 {
		w.camera.Position.X = w.Me.Position.X - 51*objects.Sign(diff.X)
	}
	if math.Abs(float64(diff.Y)) > 50 {
		w.camera.Position.Y = w.Me.Position.Y - 51*objects.Sign(diff.Y)
	}
}

func (w *Window) handleKeyboardEvent(e *sdl.KeyboardEvent, toogle bool) {
	switch e.Keysym.Sym {
	case sdl.K_F3:
		if !toogle {
			w.show_debug = !w.show_debug
		}
	default:
		w.Me.Lock()
		w.Me.HandleKeyboard(e.Keysym.Sym, toogle)
		w.Me.Unlock()
	}
}
