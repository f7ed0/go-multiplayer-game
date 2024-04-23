package animatedsprite

import (
	"encoding/json"
	"io"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Aesprite struct {
	Texture *sdl.Texture
	Frames  []*AespriteFrame `json:"frames"`
	Meta    AespriteMeta     `json:"meta"`
}

type AespriteMeta struct {
	FrameTags map[string]AespriteMetaFrameTag `json:"frameTags"`
}

type AespriteMetaFrameTag struct {
	Name string `json:"name"`
	From int    `json:"from"`
	To   int    `json:"to"`
}

type AespriteFrame struct {
	Frame    sdl.Rect `json:"frame"`
	Duration int      `json:"duration"`
}

func NewAesprite(json_path string, texture_path string, r *sdl.Renderer) (sprite *Aesprite, err error) {
	f, err := os.Open(json_path)
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
	if err != nil {
		return
	}
	surf, err := img.Load(texture_path)
	if err != nil {
		return
	}
	tex, err := r.CreateTextureFromSurface(surf)
	if err != nil {
		return
	}
	surf.Free()
	sprite.Texture = tex

	return
}

func (aes *Aesprite) GetFrame(animation string, frame int) sdl.Rect {
	v, ok := aes.Meta.FrameTags[animation]

	if ok {
		return aes.Frames[v.From+((frame/150)%(v.To-v.From+1))].Frame
	}
	return sdl.Rect{}
}

func (aes *Aesprite) GetTexture() *sdl.Texture {
	return aes.Texture
}
