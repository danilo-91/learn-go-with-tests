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

func TestAdd(t *testing.T) {
    t.Run("add word", func(t *testing.T) {
        d := Dictionary{}
        word := "test"
        want := "this is the Add test"
        d.Add(word, want)

        got, err := d.Search(word)

        if err != nil {
            t.Fatal(err.Error())
        }

        assertStrings(t, got, want)
    })

    t.Run("error on repeated word", func(t *testing.T) {
        d := Dictionary{}
        word := "test"
        def := "test already set"

        ok := d.Add(word, def)

        if !ok {
            t.Fatal("word should be added")
        }

        // Now let's make a mistake

        ok = d.Add(word, "this shouldn't be here")

        if ok {
            t.Fatal("word should not change!")
        }

        // Check definition still the same
        got, _ := d.Search(word)

        assertStrings(t, got, def)
    })
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}
