package display

import (
	"time"

	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/veandco/go-sdl2/ttf"
)

const FPS int64 = 144

func (w *Window) MainLoop() {
	w.Setup()
	var timing time.Duration = time.Duration((1000000000 / FPS)) - 590*time.Microsecond
	lg.Debug.Println(timing)
	var t time.Time = time.Now()
	w.debug.FrameCount = 0
	t2 := time.Now()
	for !w.exit {

		if time.Since(t) > 400*time.Millisecond {
			w.debug.FrameCountBuffer.Stack(w.debug.FrameCount)
			w.debug.FrameCount = 0
			t = t2
		}

		w.event()
		w.render()

		w.debug.FrameCount++
		w.debug.FrameTime = int(time.Since(t2).Microseconds())
		w.debug.FrameTimeBuffer.Stack(w.debug.FrameTime)
		time.Sleep(time.Duration(timing - time.Since(t2)))
		t2 = time.Now()
	}

}

func (w *Window) Setup() {
	var err error
	w.font, err = ttf.OpenFont("assets/fonts/Exo2-Medium.ttf", 20)
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
}
