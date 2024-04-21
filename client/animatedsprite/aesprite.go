package animatedsprite

import (
	"encoding/json"
	"io"
	"os"

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
	Frame    sdl.Rect `json:"frame"`
	Rotated  bool     `json:"rotated"`
	Trimmed  bool     `json:"trimmed"`
	Duration int      `json:"duration"`
}

func NewAesprite(path string) (sprite *Aesprite, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	var jsonb []byte
	jsonb, err = io.ReadAll(f)
	if err != nil {
		return
	}
	sprite = new(Aesprite)
	*sprite = Aesprite{}
	err = json.Unmarshal(jsonb, sprite)
	return
}
