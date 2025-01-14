package tdt

import (
	"testing"
)

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Length: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.16},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
		}
	}
}

// func TestArea(t *testing.T) {
// 	checkShapeArea := func(t testing.TB, shape Shape, want float64) {
// 		got := shape.Area()
// 		if !reflect.DeepEqual(got, want) {
// 			t.Errorf("Expect %.2f, got %.2f", want, got)
// 		}
// 	}
//
// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		want := 72.0
//
// 		checkShapeArea(t, rectangle, want)
// 	})
//
// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		want := 314.16
//
// 		checkShapeArea(t, circle, want)
// 	})
// }
