package objects

import "math"

type Point struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type Vector Point

func (p *Point) Add(other Vector) {
	p.X += other.X
	p.Y += other.Y
	p.Z += other.Z
}

func Diff(a, b Point) Vector {
	return Vector{X: b.X - a.X, Y: b.Y - a.Y, Z: b.Z - a.Z}
}

func (v Vector) N2_2D() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func Sign(x float32) float32 {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}
