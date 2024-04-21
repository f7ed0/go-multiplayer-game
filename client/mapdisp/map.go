package mapdisp

import (
	"encoding/json"
	"errors"
	"io"
	"math"
	"os"
	"sort"

	"github.com/f7ed0/go-multiplayer-game/commons/gamemap"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/veandco/go-sdl2/sdl"
)

type GameMap struct {
	Textures []Texture
	ChunksM  ChunkMap
	ChunksD  ChunkData

	gamemap.HeadLessMap
}

type Texture struct {
	Texture  *sdl.Texture
	TileSize int
}

type GameMapTextureLoader struct {
	File     string `json:"file"`
	Tilesize int    `json:"tilesize"`
}

type ChunkMap struct {
	ChunkSize int         `json:"chunksize"`
	TileSize  int         `json:"tilesize"`
	ChunkList []ChunkInfo `json:"map"`
}

type ChunkInfo struct {
	X    int    `json:"x"`
	Y    int    `json:"y"`
	Name string `json:"name"`
}

type ChunkData map[string]ChunkRaw

type ChunkRaw struct {
	UnderLayers [][]TilesRaw `json:"underlayers"`
	OverLayers  [][]TilesRaw `json:"overlayers"`
	Props       []Prop       `json:"props"`
}

type Prop struct {
	MaxY  int
	Tiles []FreeTileRaw `json:"tiles"`
}

type FreeTileRaw struct {
	TilesRaw
	X int `json:"x"`
	Y int `json:"y"`
}

type TilesRaw struct {
	Texture int `json:"texture"`
	TileX   int `json:"tile_x"`
	TileY   int `json:"tile_y"`
}

func (cr *ChunkRaw) calcProps() {
	for i := 0; i < len(cr.Props); i++ {
		cr.Props[i].MaxY = math.MinInt
		for _, tile := range cr.Props[i].Tiles {
			if tile.Y > cr.Props[i].MaxY {

				cr.Props[i].MaxY = tile.Y
			}
		}
		cr.Props[i].MaxY++
	}

}

func (c *ChunkRaw) sortProp() {
	sort.Slice(c.Props, func(i, j int) bool {
		return (c.Props)[i].MaxY < (c.Props)[j].MaxY
	})
}

func (g *GameMap) GetChunkAt(x int, y int) (res ChunkRaw, err error) {
	if g.ChunksD == nil {
		g.ChunksD = make(ChunkData)
	}
	for _, item := range g.ChunksM.ChunkList {
		if item.X == x && item.Y == y {
			v, ok := g.ChunksD[item.Name]
			if !ok {
				var f *os.File
				var bt []byte
				f, err = os.Open(g.Path + "/chunks/" + item.Name + ".json")
				if err != nil {
					return
				}
				bt, err = io.ReadAll(f)
				if err != nil {
					return
				}
				f.Close()
				err = json.Unmarshal(bt, &res)
				if err != nil {
					return
				}

				res.calcProps()
				lg.Debug.Println(res.Props)
				res.sortProp()

				lg.Debug.Println(res.Props)

				g.ChunksD[item.Name] = res
				lg.Debug.Println(g.ChunksD[item.Name].OverLayers)

				return
			}
			res = v
			return
		}
	}
	err = errors.New("no chunk found")
	return
}
