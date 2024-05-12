package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Brey")
	got := buffer.String()
	want := "Hello, Brey"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
