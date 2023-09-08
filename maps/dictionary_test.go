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
		want := ErrWordNotFound

		if err == nil {
			t.Fatal("expected error")
		}

		assertError(t, err, want)
	})

}

func TestAdd(t *testing.T) {
	t.Run("add word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		want := "this is the Add test"
		d.Add(word, want)

		got, err := d.Search(word)

		assertError(t, err, nil)
		assertStrings(t, got, want)
	})

	t.Run("error on repeated word", func(t *testing.T) {
		d := Dictionary{}
		word := "test"
		def := "test already set"

		err := d.Add(word, def)

		assertError(t, err, nil)

		// Now let's make a mistake

		err = d.Add(word, "this shouldn't be here")

		assertError(t, err, ErrWordExists)

		// Check definition still the same
		got, _ := d.Search(word)

		assertStrings(t, got, def)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is a test"
		updated := "this a updated test"
		d := Dictionary{word: def}
		d.Update(word, updated)

		got, _ := d.Search(word)
		assertStrings(t, got, updated)
	})

	t.Run("non existing word", func(t *testing.T) {
		word := "test"
		updated := "this is a updated test"
		d := Dictionary{}
		err := d.Update(word, updated)

		assertError(t, err, ErrWordDoesNotExists)
	})

}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		def := "this is a test"
		d := Dictionary{word: def}

		err := d.Delete(word)

		assertError(t, err, nil)

		_, err = d.Search(word)

		assertError(t, err, ErrWordNotFound)
	})
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
