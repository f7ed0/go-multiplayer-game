package display

import (
	"math"
	"time"

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
	w.OtherMutex.Lock()
	for i := range w.Other {
		w.Other[i].ApplyEvent(
			float32(math.Min(
				float64(delta),
				float64(time.Since(w.Other[i].LastTime))/1000,
			)),
		)
	}
	w.OtherMutex.Unlock()
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
