package interfaces

import "math"

// Shape defines Area
type Shape interface {
	Area() float64
}

// Rectangle is reactangle
type Rectangle struct {
	Length float64
	Width  float64
}

// Circle is a circle
type Circle struct {
	Radius float64
}

// Triangle is a triangle
type Triangle struct {
	Height float64
	width  float64
}

// Area computes a rectangle's area
func (rec Rectangle) Area() float64 {
	return rec.Length * rec.Width
}

// Area computes a circle's area
func (cir Circle) Area() float64 {
	return math.Pi * cir.Radius * cir.Radius
}

// Area computes a triangle's area
func (tri Triangle) Area() float64 {
	return 0.5 * tri.Height * tri.width
}
