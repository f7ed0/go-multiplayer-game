package display

import (
	"encoding/json"
	"io"
	"os"
	"sync"

	"github.com/f7ed0/go-multiplayer-game/client/handleplayer"
	"github.com/f7ed0/go-multiplayer-game/client/mapdisp"
	"github.com/f7ed0/go-multiplayer-game/commons/entity/player"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Window struct {
	self     *sdl.Window
	renderer *sdl.Renderer

	font    *ttf.Font
	exit    bool
	debug   DebugStat
	GameMap mapdisp.GameMap

	Me         handleplayer.HandledPlayer
	Other      []player.PlayerCore
	OtherMutex sync.RWMutex

	show_debug bool // TODO : move it to debug stat
}

type DebugStat struct {
	FrameCount       int
	FrameCountBuffer objects.IntBuffer
	FrameTimeBuffer  objects.IntBuffer
}

func NewWindow() (*Window, error) {
	if !subsys_initialised {
		err := init_subsystems()
		if err != nil {
			return nil, err
		}
	}
	w, r, err := sdl.CreateWindowAndRenderer(1280, 720, sdl.WINDOW_ALLOW_HIGHDPI)
	if err != nil {
		return nil, err
	}
	return &Window{
		self:     w,
		renderer: r,
		exit:     false,
		debug: DebugStat{
			FrameCountBuffer: objects.NewIntBuffer(30),
			FrameTimeBuffer:  objects.NewIntBuffer(100),
		},
		Me: handleplayer.NewHandledPlayer(),
	}, nil
}

func (w *Window) LoadMap(path string) (err error) {
	w.GameMap.Path = path
	// LOADING HITBOXES
	f, err := os.Open(path + "/hitboxes.json")
	if err != nil {
		return
	}
	res, err := io.ReadAll(f)
	if err != nil {
		return
	}
	f.Close()
	err = json.Unmarshal(res, &w.GameMap)
	if err != nil {
		return
	}
	lg.Debug.Println(w.GameMap)

	// LOADING TEXTURES
	var texs []mapdisp.GameMapTextureLoader
	f, err = os.Open(path + "/textures.json")
	if err != nil {
		return
	}
	res, err = io.ReadAll(f)
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
	w.GameMap.Textures = []mapdisp.Texture{}
	for _, item := range texs {
		surface, err = img.Load(path + "/textures/" + item.File)
		if err != nil {
			return
		}
		texture, err = w.renderer.CreateTextureFromSurface(surface)
		if err != nil {
			return
		}
		surface.Free()
		w.GameMap.Textures = append(
			w.GameMap.Textures,
			mapdisp.Texture{
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
	err = json.Unmarshal(res, &w.GameMap.ChunksM)
	if err != nil {
		return
	}

	return
}
