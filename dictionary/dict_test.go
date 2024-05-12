package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dict := Dict{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dict.Search("unknown")
		// .Error() converts error to string
		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dict := Dict{}
	dict.Add("test", "another test")
	want := "another test"
	got, err := dict.Search("test")
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
