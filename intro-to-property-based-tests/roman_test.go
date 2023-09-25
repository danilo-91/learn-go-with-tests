package roman

import (
	"testing"
)

func TestNToRoman(t *testing.T) {
    t.Run("I", func(t *testing.T) {
        got := NToRoman(1)
        want := "I"
        assertString(t, got, want)
    })

    t.Run("III", func(t *testing.T) {
        got := NToRoman(3)
        want := "III"
        assertString(t, got, want)
    })

    t.Run("IV", func(t *testing.T) {
        got := NToRoman(4)
        want := "IV"
        assertString(t, got, want)
    })
}

func assertString(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q, but wanted %q", got, want)
    }
}

