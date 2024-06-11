package mapdisp

import (
	"encoding/json"
	"io"
	"os"

	"github.com/f7ed0/go-multiplayer-game/commons/gamemap"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func LoadMap(r *sdl.Renderer, path string) (gm GameMap, err error) {
	gm.HeadLessMap, err = gamemap.LoadMap(path)
	if err != nil {
		return
	}
	// LOADING TEXTURES
	var texs []GameMapTextureLoader
	f, err := os.Open(path + "/textures.json")
	if err != nil {
		return
	}
	res, err := io.ReadAll(f)
	if err != nil {
		return
	}
	f.Close()
	err = json.Unmarshal(res, &texs)
	if err != nil {
		return
	}
	var surface *sdl.Surface
	var texture *sdl.Texture
	gm.Textures = []Texture{}
	for _, item := range texs {
		surface, err = img.Load(path + "/textures/" + item.File)
		if err != nil {
			return
		}
		texture, err = r.CreateTextureFromSurface(surface)
		if err != nil {
			return
		}
		surface.Free()
		gm.Textures = append(
			gm.Textures,
			Texture{
				Texture:  texture,
				TileSize: item.Tilesize,
			},
		)
	}

	// LOADING CHUNKS
	f, err = os.Open(path + "/chunkmap.json")
	if err != nil {
		return
	}
	res, err = io.ReadAll(f)
	if err != nil {
		return
	}
	f.Close()
	err = json.Unmarshal(res, &gm.ChunksM)
	if err != nil {
		return
	}

	return
}
