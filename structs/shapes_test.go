package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Reactangle{10.0, 10.0}
	got := rect.Perimeter()
	want := 40.0

	//  %.2f - display two decimal place float64
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Reactangle{12.0, 8.0}
		want := 96.0
		checkArea(t, rect, want)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}
