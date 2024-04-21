package display

import (
	"sync"

	"github.com/f7ed0/go-multiplayer-game/client/handleplayer"
	"github.com/f7ed0/go-multiplayer-game/client/mapdisp"
	"github.com/f7ed0/go-multiplayer-game/commons/entity/player"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
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

	width  int
	height int

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
	w, r, err := sdl.CreateWindowAndRenderer(1280, 720, sdl.WINDOW_ALLOW_HIGHDPI|sdl.WINDOW_RESIZABLE)
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
		width:  1280,
		height: 720,
		Me:     handleplayer.NewHandledPlayer(),
	}, nil
}

func (w *Window) LoadMap(path string) (err error) {
	w.GameMap, err = mapdisp.LoadMap(w.renderer, path)
	return
}
