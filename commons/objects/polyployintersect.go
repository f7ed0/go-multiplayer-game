package objects

func PolyPolyIntersect(poly1 Polygon, poly2 Polygon) bool {
	for i := 0; i < len(poly1.Points); i++ {
		for j := 0; j < len(poly2.Points); j++ {
			seg1 := Segment{
				A: poly1.Points[i],
				B: poly1.Points[(i+1)%len(poly1.Points)],
			}
			seg2 := Segment{
				A: poly2.Points[j],
				B: poly2.Points[(j+1)%len(poly2.Points)],
			}
			_, ok := SegIntersect(seg1, seg2)
			if ok {
				return true
			}
		}

	}

	return false
}
