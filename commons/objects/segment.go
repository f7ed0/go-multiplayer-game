package objects

import "math"

type Segment struct {
	A Point
	B Point
}

// Returns the normalised director vector
func (s *Segment) Director() Vector {
	p := Point{}
	if s.A == p && s.B == p {
		return Vector{}
	}
	return Diff2D(s.A, s.B).Normalized2D()
}

// Test if two segment intersect and if so return the postion of the intersection
func SegIntersect(s1, s2 Segment) (Vector, bool) {
	ba := Diff(s1.B, s1.A)
	ab := Diff(s1.A, s1.B)
	ba2 := Diff(s2.B, s2.A)
	aa := Diff(s2.A, s1.A)
	a := Cross(aa, ba2)
	b := Cross(ba, ba2)
	c := Cross(ab, aa)
	if (a < 0 && 0 < b) || (a > 0 && 0 > b) || (c < 0 && 0 < b) || (c > 0 && 0 > b) || math.Abs(float64(a)) > math.Abs(float64(b)) || math.Abs(float64(c)) > math.Abs(float64(b)) {
		return Vector{}, false
	} else {
		t := a / b
		return Vector(Sum2D(s1.A, Point(ab.Times(t)))), true
	}
}
