package animatedsprite

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Aesprite struct {
	Texture *sdl.Texture
	Frames  map[string]AespriteFrame `json:"frames"`
	meta    AespriteMeta
}

type AespriteMeta struct {
	frameTags []AespriteMetaFrameTag
}

type AespriteMetaFrameTag struct {
	Name      string `json:"name"`
	From      int    `json:"from"`
	To        int    `json:"to"`
	Direction string `json:"direction"`
}

type AespriteFrame struct {
	Frame    *sdl.Rect
	Rotated  bool
	Trimmed  bool
	Duration time.Duration
}

func NewAesprite(path string) {
	// TODO
}
