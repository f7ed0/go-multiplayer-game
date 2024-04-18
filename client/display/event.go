package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

func (w *Window) event() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch e.GetType() {
		case sdl.QUIT:
			w.exit = true
		case sdl.KEYDOWN:
			w.handleKeyboardEvent(e.(*sdl.KeyboardEvent), true)
		}
	}
}

func (w *Window) handleKeyboardEvent(e *sdl.KeyboardEvent, toogle bool) {
	switch e.Keysym.Sym {
	case sdl.K_F3:
		w.show_debug = !w.show_debug
	case sdl.K_z:
		w.camera.Position.X += 100
	default:

	}
}
