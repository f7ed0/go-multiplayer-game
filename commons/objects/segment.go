package objects

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
