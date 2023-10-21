package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("calling Inc() 3 times leaves Value() at 3", func(t *testing.T) {
		c := Counter{}
		for i := 0; i < 3; i++ {
			c.Inc()
		}
		got := c.Value()
		want := 3

		assertInt(t, got, want)
	})

	t.Run("runs safely with concurrency", func(t *testing.T) {
		c := Counter{}
		n := 1000

		var wg sync.WaitGroup
		wg.Add(n)

		for i := 0; i < n; i++ {
			go func() {
				c.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertInt(t, c.Value(), n)
	})
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %d but got %d", want, got)
	}
}
