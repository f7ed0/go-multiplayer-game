package objects

import "github.com/veandco/go-sdl2/sdl"

type Polygon struct {
	Points []Point `json:"points"`
}

func (p *Polygon) Draw(r *sdl.Renderer) {

}
