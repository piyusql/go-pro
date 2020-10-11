package boids

import "math"

type vector2D struct {
	x float64
	y float64
}

func (v1 vector2D) X() int {
	return int(v1.x)
}
func (v1 vector2D) Y() int {
	return int(v1.y)
}

func (v1 vector2D) Add(v2 vector2D) vector2D {
	return vector2D{v1.x + v2.x, v1.y + v2.y}
}

func (v1 vector2D) Subtract(v2 vector2D) vector2D {
	return vector2D{v1.x - v2.x, v1.y - v2.y}
}

func (v1 vector2D) Distance(v2 vector2D) float64 {
	return math.Hypot(v1.x-v2.x, v1.y-v2.y)
}

func (v1 vector2D) Limit(lower, upper float64) vector2D {
	return vector2D{math.Min(math.Max(lower, v1.x), upper),
		math.Min(math.Max(lower, v1.y), upper)}
}

func (v1 vector2D) AddV(val float64) vector2D {
	return vector2D{v1.x + val, v1.y + val}
}

func (v1 vector2D) SubtractV(val float64) vector2D {
	return vector2D{v1.x - val, v1.y - val}
}

func (v1 vector2D) MultiplyV(val float64) vector2D {
	return vector2D{v1.x * val, v1.y * val}
}

func (v1 vector2D) DivisionV(val float64) vector2D {
	return vector2D{v1.x / val, v1.y / val}
}
