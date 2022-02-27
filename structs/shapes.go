package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Triangle struct {
	Base   float64
	height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.height / 2
}

func Perimeter(r Rectangle) float64 {
	return (r.Width + r.Height) * 2
}
