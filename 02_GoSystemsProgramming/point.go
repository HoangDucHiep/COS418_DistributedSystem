package main

import "math"

type Point struct {
	X, Y float64
}

// Implement a method that find Hypotenuse distance between one Point and another
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// standard function
func Distance(p Point, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
