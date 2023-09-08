package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "this is a test"}

	got := Search(dict, "test")
	want := "this is a test"
    assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q, but wanted %q", got, want)
    }
}
