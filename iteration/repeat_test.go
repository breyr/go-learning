package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("repeat default amount", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertEquals(t, got, want)
	})
	t.Run("repeat custom amount", func(t *testing.T) {
		got := Repeat("a", 10)
		want := "aaaaaaaaaa"
		assertEquals(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func assertEquals(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("expected %q but got %q", got, want)
	}
}

func ExampleRepeat() {
	str := Repeat("a", 5)
	fmt.Println(str)
	// Output: aaaaa
}
