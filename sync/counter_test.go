package counter

import "testing"

func TestCounter(t *testing.T) {
	t.Run("calling Inc() 3 times leaves Value() at 3", func(t *testing.T) {
		c := Counter{}
		for i := 0; i < 3; i++ {
			c.Inc()
		}
		want := 3

        assertCounter(t, c, want)
	})
}

func assertCounter(t testing.TB, got Counter, want int) {
    t.Helper()
    if got.Value() != want {
        t.Errorf("wanted %d but got %d", want, got.Value())
    }
}
