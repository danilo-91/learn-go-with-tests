package main

import (
    "testing"
    "bytes"
    "reflect"
)

type SpyCountdownSleeper struct {
    Calls []string
}

func (s *SpyCountdownSleeper) Sleep() {
    s.Calls = append(s.Calls, "sleep")
}

func (s *SpyCountdownSleeper) Write(b []byte) (n int, err error) {
    s.Calls = append(s.Calls, "write")
    return
}

func TestCountdown(t *testing.T) {
    t.Run("text", func (t *testing.T) {
        b := &bytes.Buffer{}
        sl := &SpyCountdownSleeper{}

        Countdown(b, sl)

        got := b.String()
        want := `3
2
1
Go!
`

        if got != want {
            t.Errorf("wanted %q, but got %q", want, got)
        }
    })

    t.Run("3 write 3 sleep", func (t *testing.T) {
        sl := &SpyCountdownSleeper{}

        Countdown(sl, sl)

        want := []string{
            "write",
            "sleep",
            "write",
            "sleep",
            "write",
            "sleep",
            "write",
        }


        if !reflect.DeepEqual(want, sl.Calls) {
            t.Errorf("wanted %q, but got %q", want, sl.Calls)
        }
    })
}
