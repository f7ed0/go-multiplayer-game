package objects

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
