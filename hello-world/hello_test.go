package main

import "testing"

func TestHello(t *testing.T) {
    t.Run("saying hello to people", func (t *testing.T) {
        got := Hello("Danilo", "en")
        want := "Hello, Danilo!"
        assertOutput(t, got, want)
    })

    t.Run("say 'Hello, World!' when empty string is provided", func (t *testing.T) {
        got := Hello("", "en")
        want := "Hello, World!"
        assertOutput(t, got, want)
    })

    t.Run("say Hello in spanish", func (t *testing.T) {
        got := Hello("Danilo", "es")
        want := "Hola, Danilo!"
        assertOutput(t, got, want)
    })

    t.Run("say hello in french", func (t *testing.T) {
        got := Hello("Danilo", "fr")
        want := "Bonjour, Danilo!"
        assertOutput(t, got, want)
    })

    t.Run("say hello but in italian", func (t *testing.T) {
        got := Hello("Danilo", "it")
        want := "Ciao, Danilo!"
        assertOutput(t, got, want)
    })
}

func assertOutput(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q wanted %q", got, want)
    }
}
