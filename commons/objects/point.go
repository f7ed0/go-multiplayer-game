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

func Sum2D(a, b Point) Point {
	return Point{X: a.X + b.X, Y: a.Y + b.Y, Z: 0}
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

// CALC THE VERCTOR BA
func Diff2D(a, b Point) Vector {
	return Vector{X: b.X - a.X, Y: b.X - a.X}
}

func Div2D(v Vector, amount float32) Vector {
	if amount != 0 {
		return Vector{X: v.X / amount, Y: v.Y / amount}
	}
	return v
}

func (v Vector) Normalized2D() Vector {

	return Div2D(v, v.N2_2D())

}
