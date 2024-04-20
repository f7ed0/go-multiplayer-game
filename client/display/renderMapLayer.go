package display

import (
	"github.com/f7ed0/go-multiplayer-game/client/mapdisp"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/veandco/go-sdl2/sdl"
)

func (w *Window) RenderMapLayer(layers [][]mapdisp.TilesRaw, offset_x, offset_y int) {
	r_o_x := offset_x * w.GameMap.ChunksM.ChunkSize * w.GameMap.ChunksM.TileSize
	r_o_y := offset_y * w.GameMap.ChunksM.ChunkSize * w.GameMap.ChunksM.TileSize
	for _, layer := range layers {
		for i := 0; i < w.GameMap.ChunksM.ChunkSize; i++ {
			for j := 0; j < w.GameMap.ChunksM.ChunkSize; j++ {
				if j*w.GameMap.ChunksM.ChunkSize+i > len(layer)-1 {
					continue
				}
				if layer[j*w.GameMap.ChunksM.ChunkSize+i].Texture == -1 {
					continue
				}
				x := int32(layer[j*w.GameMap.ChunksM.ChunkSize+i].TileX * w.GameMap.Textures[layer[j*w.GameMap.ChunksM.ChunkSize+i].Texture].TileSize)
				y := int32(layer[j*w.GameMap.ChunksM.ChunkSize+i].TileY * w.GameMap.Textures[layer[j*w.GameMap.ChunksM.ChunkSize+i].Texture].TileSize)
				width := int32(w.GameMap.Textures[layer[j*w.GameMap.ChunksM.ChunkSize+i].Texture].TileSize)
				height := int32(w.GameMap.Textures[layer[j*w.GameMap.ChunksM.ChunkSize+i].Texture].TileSize)
				to := w.applyOffset(objects.Point{X: float32(i*w.GameMap.ChunksM.TileSize + r_o_x), Y: float32(j*w.GameMap.ChunksM.TileSize + r_o_y)})
				to_w := w.applyOffsetF32(float32(w.GameMap.ChunksM.TileSize))
				to_h := w.applyOffsetF32(float32(w.GameMap.ChunksM.TileSize))
				//lg.Debug.Println(x, y, width, height)
				err := w.renderer.CopyF(
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

func (w *Window) renderProp(prop mapdisp.Prop, r_o_x, r_o_y int) {

	for _, tile := range prop.Tiles {
		x := int32(tile.TileX * w.GameMap.Textures[tile.Texture].TileSize)
		y := int32(tile.TileY * w.GameMap.Textures[tile.Texture].TileSize)
		width := int32(w.GameMap.Textures[tile.Texture].TileSize)
		height := int32(w.GameMap.Textures[tile.Texture].TileSize)
		to := w.applyOffset(objects.Point{X: float32(tile.X*w.GameMap.ChunksM.TileSize + r_o_x), Y: float32(tile.Y*w.GameMap.ChunksM.TileSize + r_o_y)})
		to_w := w.applyOffsetF32(float32(w.GameMap.ChunksM.TileSize))
		to_h := w.applyOffsetF32(float32(w.GameMap.ChunksM.TileSize))
		err := w.renderer.CopyF(
			w.GameMap.Textures[tile.Texture].Texture,
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
