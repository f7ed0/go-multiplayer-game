package display

import (
	"fmt"

	"github.com/f7ed0/go-multiplayer-game/client/display/sdlplus"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/veandco/go-sdl2/sdl"
)

func (w *Window) render(delta int) {
	w.renderer.SetDrawColor(0, 0, 0, 255)
	w.renderer.Clear()
	w.mapUnderLayer()
	w.PlayerLayer(delta)
	w.mapOverLayer()
	if w.show_debug {
		w.hiboxLayer()
		w.debugLayer()
	}
	w.renderer.Present()
	w.renderer.Flush()
}

func (w *Window) mapOverLayer() {
	chunk, err := w.GameMap.GetChunkAt(0, 0)
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
	w.RenderMapLayer(chunk.OverLayers, 0, 0)

}

func (w *Window) mapUnderLayer() {
	for _, item := range w.GameMap.ChunksM.ChunkList {
		chunk, err := w.GameMap.GetChunkAt(item.X, item.Y)
		if err != nil {
			lg.Error.Fatalln(err.Error())
		}
		w.RenderMapLayer(chunk.UnderLayers, item.X, item.Y)
	}

}

func (w *Window) debugLayer() {
	fps_text := fmt.Sprintf("%.0f FPS", w.debug.FrameCountBuffer.GetMean()*2.5)
	sdlplus.RenderText(w.renderer, w.font, fps_text, YELLOW, 10, objects.Point{X: 0, Y: 0})

	ft_text := fmt.Sprintf("%.0f µs FT", w.debug.FrameTimeBuffer.GetMean())
	sdlplus.RenderText(w.renderer, w.font, ft_text, ORANGE, 10, objects.Point{X: 0, Y: 21})

	size_text := fmt.Sprintf("res : %vx%v", w.width, w.height)
	sdlplus.RenderText(w.renderer, w.font, size_text, ORANGE, 10, objects.Point{X: 0, Y: 42})
}

var (
	YELLOW = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	ORANGE = sdl.Color{R: 255, G: 200, B: 0, A: 255}
)

func (w *Window) applyOffset(p objects.Point) objects.Point {
	return w.Me.Camera.ApplyOffset(p, w.width, w.height)
}

func (w *Window) applyOffsetF32(value float32) float32 {
	return w.Me.Camera.ApplyOffsetF32(value)
}
