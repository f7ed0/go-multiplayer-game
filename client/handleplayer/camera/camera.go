package camera

import (
	"github.com/f7ed0/go-multiplayer-game/commons/objects"
)

type Camera struct {
	Position objects.Point
}

func (cam Camera) ApplyOffset(p objects.Point, width, height int) objects.Point {
	return objects.Point{
		X: (p.X-cam.Position.X)*cam.Position.Z + float32(width)/2,
		Y: (p.Y-cam.Position.Y)*cam.Position.Z + float32(height)/2,
		Z: p.Z,
	}
}

func (cam Camera) ApplyOffsetF32(value float32) float32 {
	return value * cam.Position.Z
}
