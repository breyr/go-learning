package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Reactangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	//  %.2f - display two decimal place float64
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	rect := Reactangle{12.0, 8.0}
	got := Area(rect)
	want := 96.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
