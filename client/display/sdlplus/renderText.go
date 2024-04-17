package sdlplus

import (
	"github.com/f7ed0/go-multiplayer-game/commons/lg"
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func RenderText(r *sdl.Renderer, font *ttf.Font, text string, color sdl.Color, lettersize float32, origin objects.Point) {
	surf, err := font.RenderUTF8Blended(text, color)
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
	tex, err := r.CreateTextureFromSurface(surf)
	if err != nil {
		lg.Error.Fatalln(err.Error())
	}
	surf.Free()
	r.CopyF(tex, nil, &sdl.FRect{X: float32(origin.X), Y: float32(origin.Y), W: float32(len(text)) * lettersize, H: lettersize * 2})
	tex.Destroy()
}
