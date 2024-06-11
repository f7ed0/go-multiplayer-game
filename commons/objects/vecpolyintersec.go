package objects

// Check if a verctor intersect with a polygon
func VectPolyIntersect(start Point, direction Vector, poly Polygon, radius float64) bool {

	tankDir := Segment{
		A: start,
		B: start.Add(direction),
	}

	for i := 0; i < len(poly.Points); i++ {
		seg := Segment{
			A: poly.Points[i],
			B: poly.Points[(i+1)%len(poly.Points)],
		}
		_, ok := SegIntersect(tankDir, seg)
		if ok {
			return true
		}
	}

	return false
}
