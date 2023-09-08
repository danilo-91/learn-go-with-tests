package dictionary

import "testing"

func TestSearch(t *testing.T) {

	t.Run("word that exists", func(t *testing.T) {
		d := Dictionary{"test": "this is a test"}
		got, _ := d.Search("test")
		want := "this is a test"
		assertStrings(t, got, want)
	})

	t.Run("word that does not exist", func(t *testing.T) {
		d := Dictionary{}
		_, err := d.Search("snorlax")
		want := "Error: word \"snorlax\" not found!"

		if err == nil {
			t.Fatal("expected error")
		}

		assertStrings(t, err.Error(), want)
	})

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}
