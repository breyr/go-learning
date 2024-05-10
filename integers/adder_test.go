package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	got := Add(2, 2)
	want := 4

	// %d to print digit
	if got != want {
		t.Errorf("expected '%d' but got '%d'", got, want)
	}
}

// Testable functions added to godoc
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
