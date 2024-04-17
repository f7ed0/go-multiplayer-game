package display

import (
	"encoding/json"
	"io"
	"os"

	"github.com/f7ed0/go-multiplayer-game/commons/gamemap"
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Window struct {
	Self     *sdl.Window
	Renderer *sdl.Renderer

	font    *ttf.Font
	exit    bool
	debug   DebugStat
	GameMap gamemap.GameMap

	show_debug bool
}

type DebugStat struct {
	FrameCount       int
	FrameCountBuffer objects.IntBuffer
	FrameTime        int
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
		Self:     w,
		Renderer: r,
		exit:     false,
		debug: DebugStat{
			FrameCountBuffer: objects.NewIntBuffer(100),
			FrameTimeBuffer:  objects.NewIntBuffer(100),
		},
	}, nil
}

func (w *Window) LoadMap(path string) (err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	res, err := io.ReadAll(f)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &w.GameMap)
	if err != nil {
		return
	}
	lg.Debug.Println(w.GameMap)
	return
}
