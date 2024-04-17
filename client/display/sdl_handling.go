package display

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var subsys_initialised = false

func init_subsystems() (err error) {
	err = sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		return
	}
	err = ttf.Init()
	if err != nil {
		return
	}
	subsys_initialised = true
	return
}
