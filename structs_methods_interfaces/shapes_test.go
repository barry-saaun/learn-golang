package structs

import (
	"reflect"
	"testing"
)

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	checkShapeArea := func(t testing.TB, shape Shape, want float64) {
		got := shape.Area()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expect %.2f, got %.2f", want, got)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0

		checkShapeArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.16

		checkShapeArea(t, circle, want)
	})
}
