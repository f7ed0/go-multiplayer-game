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
			switch e.(*sdl.KeyboardEvent).Keysym.Scancode {
			case sdl.SCANCODE_F3:
				w.show_debug = !w.show_debug
			}
		}
	}
}
