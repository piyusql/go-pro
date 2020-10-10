package main

type vector2D struct {
	x float64
	y float64
}

func (v1 vector2D) Add(v2 vector2D) vector2D {
	return vector2D{v1.x + v2.x, v1.y + v2.y}
}
