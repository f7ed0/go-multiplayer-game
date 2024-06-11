package objects

type Polygon struct {
	Points []Point `json:"points"`
}

func (poly Polygon) OffsetPolygon(p Point) Polygon {
	ret := Polygon{}
	for _, point := range poly.Points {
		ret.Points = append(ret.Points, Sum2D(point, p))
	}
	return ret
}
