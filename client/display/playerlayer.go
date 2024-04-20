package display

import (
	"github.com/f7ed0/go-multiplayer-game/commons/entity"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/veandco/go-sdl2/sdl"
)

func (w *Window) PlayerLayer() {
	for _, item := range w.GameMap.ChunksM.ChunkList {
		chunk, err := w.GameMap.GetChunkAt(item.X, item.Y)
		r_o_x := item.X * w.GameMap.ChunksM.ChunkSize * w.GameMap.ChunksM.TileSize
		r_o_y := item.Y * w.GameMap.ChunksM.ChunkSize * w.GameMap.ChunksM.TileSize
		if err != nil {
			lg.Error.Fatalln(err.Error())
		}
		var a []entity.Entity = []entity.Entity{}
		w.Me.RLock()
		if w.Me.Position.X > float32(r_o_x) && item.X < (r_o_x+w.GameMap.ChunksM.ChunkSize*w.GameMap.ChunksM.TileSize) {
			if w.Me.Position.Y > float32(r_o_y) && item.Y < (r_o_y+w.GameMap.ChunksM.ChunkSize*w.GameMap.ChunksM.TileSize) {

				a = append(a, w.Me.Entity)

			}
		}
		w.Me.RUnlock()
		w.OtherMutex.RLock()
		for _, otherplayer := range w.Other {
			if otherplayer.Position.X > float32(r_o_x) && item.X < (r_o_x+w.GameMap.ChunksM.ChunkSize*w.GameMap.ChunksM.TileSize) {
				if otherplayer.Position.Y > float32(r_o_y) && item.Y < (r_o_y+w.GameMap.ChunksM.ChunkSize*w.GameMap.ChunksM.TileSize) {

					a = append(a, otherplayer.Entity)

				}
			}
		}
		w.OtherMutex.RUnlock()
		var i, j int
		for i < len(chunk.Props) && j < len(a) {
			if (chunk.Props[i].MaxY*w.GameMap.ChunksM.TileSize)+r_o_y < int(a[j].Position.Y)+w.GameMap.ChunksM.TileSize {
				w.renderProp(chunk.Props[i], r_o_x, r_o_y)
				i++
			} else {
				offseted := w.applyOffset(a[j].Position)
				height := w.applyOffsetF32(float32(w.GameMap.ChunksM.TileSize))
				width := height / 1.5
				w.renderer.SetDrawColor(255, 100, 100, 255)
				w.renderer.FillRectF(&sdl.FRect{
					X: offseted.X,
					Y: offseted.Y,
					W: width,
					H: height,
				})
				j++
			}
		}
		for i < len(chunk.Props) {
			w.renderProp(chunk.Props[i], r_o_x, r_o_y)
			i++
		}
		for j < len(a) {
			offseted := w.applyOffset(a[j].Position)
			height := w.applyOffsetF32(float32(w.GameMap.ChunksM.TileSize))
			width := height / 1.5
			w.renderer.SetDrawColor(255, 100, 100, 255)
			w.renderer.FillRectF(&sdl.FRect{
				X: offseted.X,
				Y: offseted.Y,
				W: width,
				H: height,
			})
			j++
		}
	}
}
