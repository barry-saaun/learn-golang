package tdt

import (
	"math"
)

type Rectangle struct {
	Width  float64
	Length float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Length
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	area := math.Pi * math.Pow(c.Radius, 2)
	return math.Round(area*100) / 100
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
