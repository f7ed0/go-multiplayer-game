package animatedsprite

import "github.com/veandco/go-sdl2/sdl"

type AnimatedSptite interface {
	GetFrame(animation string, frame int) sdl.Rect
	GetTexture() *sdl.Texture
}
