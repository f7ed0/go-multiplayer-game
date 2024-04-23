package display

import (
	"sort"

	"github.com/f7ed0/go-multiplayer-game/client/display/drawableentity"
	"github.com/f7ed0/go-multiplayer-game/client/handleplayer"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
)

func (w *Window) PlayerLayer(delta int) {
	drawableentity.FrameCounter += delta
	for _, item := range w.GameMap.ChunksM.ChunkList {
		chunk, err := w.GameMap.GetChunkAt(item.X, item.Y)
		r_o_x := item.X * w.GameMap.ChunksM.ChunkSize * w.GameMap.ChunksM.TileSize
		r_o_y := item.Y * w.GameMap.ChunksM.ChunkSize * w.GameMap.ChunksM.TileSize
		if err != nil {
			lg.Error.Fatalln(err.Error())
		}
		var a []*handleplayer.DispPlayer = []*handleplayer.DispPlayer{}
		w.Me.RLock()
		if inChunk(r_o_x, r_o_y, w.GameMap.ChunksM.ChunkSize, w.GameMap.ChunksM.TileSize, w.Me.Position, 32) {
			a = append(
				a,
				&w.Me.DispPlayer,
			)
		}
		w.Me.RUnlock()
		w.OtherMutex.RLock()
		for _, otherplayer := range w.Other {
			if inChunk(r_o_x, r_o_y, w.GameMap.ChunksM.ChunkSize, w.GameMap.ChunksM.TileSize, otherplayer.Position, 32) && otherplayer.Here {
				a = append(
					a,
					otherplayer,
				)
			}
		}
		sort.Slice(a, func(i2, j2 int) bool {
			return a[i2].Entity.Position.Y < a[j2].Entity.Position.Y
		})
		w.OtherMutex.RUnlock()
		var i, j int
		for i < len(chunk.Props) && j < len(a) {
			if (chunk.Props[i].MaxY*w.GameMap.ChunksM.TileSize)+r_o_y < int(a[j].Position.Y)+w.GameMap.ChunksM.TileSize {
				w.renderProp(chunk.Props[i], r_o_x, r_o_y)
				i++
			} else {
				err := drawableentity.Draw(a[j], w.renderer, w.Me.Camera, w.PlayerModels, delta)
				if err != nil {
					lg.Error.Fatalln(err.Error())
				}
				j++
			}
		}
		for i < len(chunk.Props) {
			w.renderProp(chunk.Props[i], r_o_x, r_o_y)
			i++
		}
		for j < len(a) {
			err := drawableentity.Draw(a[j], w.renderer, w.Me.Camera, w.PlayerModels, delta)
			if err != nil {
				lg.Error.Fatalln(err.Error())
			}
			j++
		}
	}
}

func inChunk(r_o_x, r_o_y, chunksize, tilesize int, posE objects.Point, height float32) bool {
	a := posE.X > float32(r_o_x) && posE.X < float32(r_o_x+chunksize*tilesize)
	b := posE.Y+height > float32(r_o_y) && posE.Y+height < float32(r_o_y+chunksize*tilesize)
	return a && b
}
