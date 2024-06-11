package display

import (
	"sync"

	"github.com/f7ed0/go-multiplayer-game/client/animatedsprite"
	"github.com/f7ed0/go-multiplayer-game/client/display/drawableentity"
	"github.com/f7ed0/go-multiplayer-game/client/handleplayer"
	"github.com/f7ed0/go-multiplayer-game/client/mapdisp"
	"github.com/f7ed0/go-multiplayer-game/commons/entity"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/f7ed0/golog/lg"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type Window struct {
	// Window structure contains everything the client need to run the game
	self     *sdl.Window
	renderer *sdl.Renderer

	font    *ttf.Font
	cursors []*sdl.Cursor
	exit    bool
	debug   DebugStat
	GameMap mapdisp.GameMap

	Me           handleplayer.HandledPlayer
	Other        map[string]*handleplayer.DispPlayer
	PlayerModels map[entity.EntityType]drawableentity.EntityModel
	OtherMutex   sync.RWMutex

	width  int
	height int

	show_debug bool // TODO : move it to debug stat
}

type DebugStat struct {
	// Substructure of Window used for metrics data
	FrameCount       int
	FrameCountBuffer objects.IntBuffer
	FrameTimeBuffer  objects.IntBuffer
}

// Create a new sdl Window if all needed subsystems are initialised
func NewWindow() (*Window, error) {
	if !subsys_initialised {
		err := init_subsystems()
		if err != nil {
			return nil, err
		}
	}
	w, r, err := sdl.CreateWindowAndRenderer(1280, 720, sdl.WINDOW_ALLOW_HIGHDPI|sdl.WINDOW_RESIZABLE)
	w.SetTitle("GO MMO")
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
		width:        1280,
		height:       720,
		Me:           handleplayer.NewHandledPlayer(),
		PlayerModels: make(map[entity.EntityType]drawableentity.EntityModel),
		Other:        make(map[string]*handleplayer.DispPlayer),
	}, nil
}

// Load map files
func (w *Window) LoadMap(path string) (err error) {
	w.GameMap, err = mapdisp.LoadMap(w.renderer, path)
	return
}

// Load Player Sprite into SDL_Textures
func (w *Window) LoadPlayerModels() (err error) {
	kwalk, err := animatedsprite.NewAesprite("./assets/entity/knight/Knight-1.json", "./assets/entity/knight/Knight-1.png", w.renderer)
	if err != nil {
		return
	}
	kidle, err := animatedsprite.NewAesprite("./assets/entity/knight/Knight-1_idle.json", "./assets/entity/knight/Knight-1_idle.png", w.renderer)
	if err != nil {
		return
	}
	w.PlayerModels[entity.KNIGHT] = drawableentity.EntityModel{
		States: map[string]animatedsprite.AnimatedSptite{
			"walking": kwalk,
			"idling":  kidle,
		},
	}
	lg.Debug.Println(kwalk, kidle)
	//lg.Debug.Println(kwalk.GetFrame())
	return
}

// Load cursors PNG into SDL with SDL_CreateCursor
func (w *Window) LoadCursors() (err error) {
	surf, err := img.Load("./assets/cursors/36x36px/Cursor Default Friends.png")
	if err != nil {
		return
	}
	w.cursors = append(w.cursors, sdl.CreateColorCursor(surf, 0, 0))
	sdl.SetCursor(w.cursors[len(w.cursors)-1])
	surf.Free()
	return
}
