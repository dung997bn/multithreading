package main

import "math"

//Vector2D type
type Vector2D struct {
	x float64
	y float64
}

//Add func
func (v1 Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{v1.x + v2.x, v1.y + v2.y}
}

//Substract func
func (v1 Vector2D) Substract(v2 Vector2D) Vector2D {
	return Vector2D{v1.x - v2.x, v1.y - v2.y}
}

//Multiply func
func (v1 Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{v1.x * v2.x, v1.y * v2.y}
}

//Division func
func (v1 Vector2D) Division(d float64) Vector2D {
	return Vector2D{v1.x / d, v1.y / d}
}

//AddV func
func (v1 Vector2D) AddV(d float64) Vector2D {
	return Vector2D{v1.x + d, v1.y + d}
}

//MultiplyV func
func (v1 Vector2D) MultiplyV(d float64) Vector2D {
	return Vector2D{v1.x * d, v1.y * d}
}

func (v1 Vector2D) limit(lower, upper float64) Vector2D {
	return Vector2D{
		math.Min(math.Max(v1.x, lower), upper),
		math.Min(math.Max(v1.y, lower), upper),
	}
}

//Distance func
func (v1 Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}