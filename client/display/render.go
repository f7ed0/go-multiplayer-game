package display

import (
	"fmt"

	"github.com/f7ed0/go-multiplayer-game/client/display/sdlplus"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/veandco/go-sdl2/sdl"
)

func (w *Window) render() {
	w.Renderer.SetDrawColor(0, 0, 0, 255)
	w.Renderer.Clear()
	w.baseLayer()
	if w.show_debug {
		w.debugLayer()
	}
	w.Renderer.Present()
	w.Renderer.Flush()
}

func (w *Window) baseLayer() {
	w.Renderer.SetDrawColor(255, 0, 0, 255)
	for _, poly := range w.GameMap.Walls {
		w.DrawPolygon(poly)
	}
	w.Renderer.SetDrawColor(0, 0, 255, 255)
	for _, poly := range w.GameMap.Holes {
		w.DrawPolygon(poly)
	}

}

func (w *Window) debugLayer() {
	fps_text := fmt.Sprintf("%.0f FPS", w.debug.FrameCountBuffer.GetMean()*2.5)
	sdlplus.RenderText(w.Renderer, w.font, fps_text, YELLOW, 10, objects.Point{X: 0, Y: 0})

	ft_text := fmt.Sprintf("%.0f µs FT", w.debug.FrameTimeBuffer.GetMean())
	sdlplus.RenderText(w.Renderer, w.font, ft_text, ORANGE, 10, objects.Point{X: 0, Y: 21})

}

var (
	YELLOW = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	ORANGE = sdl.Color{R: 255, G: 200, B: 0, A: 255}
)

func (w *Window) DrawPolygon(p objects.Polygon) {
	for i := 0; i < len(p.Points)-1; i++ {
		w.Renderer.DrawLineF(p.Points[i].X, p.Points[i].Y, p.Points[i+1].X, p.Points[i+1].Y)
	}
	w.Renderer.DrawLineF(p.Points[len(p.Points)-1].X, p.Points[len(p.Points)-1].Y, p.Points[0].X, p.Points[0].Y)
}
