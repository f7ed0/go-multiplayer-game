package display

import (
	"math"
	"time"

	"github.com/f7ed0/go-multiplayer-game/commons/lg"
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
		case sdl.WINDOWEVENT:
			w.handleWindowEvent(e.(*sdl.WindowEvent))
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
				float64(time.Since(w.Other[i].LastTime).Milliseconds())/1000,
			)),
		)
	}

	w.OtherMutex.Unlock()
}

func (w *Window) handleKeyboardEvent(e *sdl.KeyboardEvent, toggle bool) {
	switch e.Keysym.Sym {
	case sdl.K_F3:
		if !toggle {
			w.show_debug = !w.show_debug
		}
	case sdl.K_F11:
		if !toggle {
			w.ToggleFullScreen(toggle)
		}
	default:
		w.Me.Lock()
		w.Me.HandleKeyboard(e.Keysym.Sym, toggle)
		w.Me.Unlock()
	}
}

func (w *Window) handleWindowEvent(e *sdl.WindowEvent) {
	switch e.Event {
	case sdl.WINDOWEVENT_SIZE_CHANGED:
		w.width = int(e.Data1)
		w.height = int(e.Data2)
		lg.Info.Println("Window Resized", e.Data1, e.Data2)
		w.Me.Camera.Position.Z = float32(math.Max(float64(w.width)/1280, float64(w.height)/720)) * 2
	}

}

func (w *Window) ToggleFullScreen(state bool) {
	if w.self.GetFlags()&sdl.WINDOW_FULLSCREEN_DESKTOP > 0 {
		w.self.SetFullscreen(0)
	} else {
		w.self.SetFullscreen(sdl.WINDOW_FULLSCREEN_DESKTOP)
	}
}
