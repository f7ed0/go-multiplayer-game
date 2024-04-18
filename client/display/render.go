package display

import (
	"fmt"

	"github.com/f7ed0/go-multiplayer-game/client/display/sdlplus"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/veandco/go-sdl2/sdl"
)

func (w *Window) render() {
	offset := w.calculateOffset()
	w.renderer.SetDrawColor(0, 0, 0, 255)
	w.renderer.Clear()
	w.baseLayer(offset)
	if w.show_debug {
		w.hiboxLayer(offset)
		w.debugLayer(offset)
	}
	w.renderer.Present()
	w.renderer.Flush()
}

func (w *Window) baseLayer(offset objects.Point) {
	chunk, err := w.GameMap.GetChunkAt(0, 0)
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
	for _, layer := range chunk {
		for i := 0; i < w.GameMap.ChunksM.ChunkSize; i++ {
			for j := 0; j < w.GameMap.ChunksM.ChunkSize; j++ {
				if j*w.GameMap.ChunksM.ChunkSize+i > len(layer)-1 {
					continue
				}
				x := int32(layer[j*w.GameMap.ChunksM.ChunkSize+i].TileX * w.GameMap.Textures[layer[j*w.GameMap.ChunksM.ChunkSize+i].Texture].TileSize)
				y := int32(layer[j*w.GameMap.ChunksM.ChunkSize+i].TileY * w.GameMap.Textures[layer[j*w.GameMap.ChunksM.ChunkSize+i].Texture].TileSize)
				width := int32(w.GameMap.Textures[layer[j*w.GameMap.ChunksM.ChunkSize+i].Texture].TileSize)
				height := int32(w.GameMap.Textures[layer[j*w.GameMap.ChunksM.ChunkSize+i].Texture].TileSize)
				to := applyOffset(objects.Point{X: float32(i * w.GameMap.ChunksM.TileSize), Y: float32(j * w.GameMap.ChunksM.TileSize)}, offset)
				to_w := applyOffsetF32(float32(w.GameMap.ChunksM.TileSize), offset)
				to_h := applyOffsetF32(float32(w.GameMap.ChunksM.TileSize), offset)
				lg.Debug.Println(x, y, width, height)
				err = w.renderer.CopyF(
					w.GameMap.Textures[layer[j*w.GameMap.ChunksM.ChunkSize+i].Texture].Texture,
					&sdl.Rect{
						X: x,
						Y: y,
						W: width,
						H: height,
					},
					&sdl.FRect{
						X: to.X,
						Y: to.Y,
						W: to_w,
						H: to_h,
					},
				)
				if err != nil {
					lg.Error.Fatalln(err.Error())
				}
			}
		}
	}
}

func (w *Window) hiboxLayer(offset objects.Point) {
	w.renderer.SetDrawColor(255, 0, 0, 255)
	for _, poly := range w.GameMap.Walls {
		w.DrawPolygon(poly, offset)
	}
	w.renderer.SetDrawColor(0, 0, 255, 255)
	for _, poly := range w.GameMap.Holes {
		w.DrawPolygon(poly, offset)
	}

}

func (w *Window) debugLayer(offset objects.Point) {
	fps_text := fmt.Sprintf("%.0f FPS", w.debug.FrameCountBuffer.GetMean()*2.5)
	sdlplus.RenderText(w.renderer, w.font, fps_text, YELLOW, 10, objects.Point{X: 0, Y: 0})

	ft_text := fmt.Sprintf("%.0f µs FT", w.debug.FrameTimeBuffer.GetMean())
	sdlplus.RenderText(w.renderer, w.font, ft_text, ORANGE, 10, objects.Point{X: 0, Y: 21})

}

var (
	YELLOW = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	ORANGE = sdl.Color{R: 255, G: 200, B: 0, A: 255}
)

func (w *Window) DrawPolygon(p objects.Polygon, offset objects.Point) {
	for i := 0; i < len(p.Points)-1; i++ {
		pi := applyOffset(p.Points[i], offset)
		pi1 := applyOffset(p.Points[i+1], offset)
		//lg.Verbose.Println(p.Points[i], offset, pi)
		w.renderer.DrawLineF(pi.X, pi.Y, pi1.X, pi1.Y)
	}
	pl := applyOffset(p.Points[len(p.Points)-1], offset)
	pf := applyOffset(p.Points[0], offset)
	w.renderer.DrawLineF(pl.X, pl.Y, pf.X, pf.Y)
}

func (w *Window) calculateOffset() objects.Point {
	width, height := w.self.GetSize()
	lg.Debug.Println(width, height)
	return objects.Point{
		X: float32(width)/2 - w.camera.Position.X,
		Y: float32(height)/2 - w.camera.Position.Y,
		Z: w.camera.Position.Z,
	}
}

func applyOffset(p, offset objects.Point) objects.Point {
	return objects.Point{
		X: (p.X + offset.X) * offset.Z,
		Y: (p.Y + offset.Y) * offset.Z,
		Z: p.Z,
	}
}

func applyOffsetF32(value float32, offset objects.Point) float32 {
	return value * offset.Z
}
