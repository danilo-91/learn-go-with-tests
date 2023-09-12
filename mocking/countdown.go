package main

import (
    "io"
    "os"
    "fmt"
    "time"
)

const (
    lastLine = "Go!"
    countdownStart = 3
)

type Sleeper interface {
    Sleep() 
}

type ConfigurableSleeper struct {
    d time.Duration
    sleep func(time.Duration)
}

func (sl *ConfigurableSleeper) Sleep() {
    sl.sleep(sl.d)
}

func Countdown(w io.Writer, sl Sleeper) {
    for i := countdownStart; i > 0; i-- {
        fmt.Fprintln(w, i)
        sl.Sleep()
    }
    fmt.Fprintln(w, lastLine)
}

func main() {
    sl := &ConfigurableSleeper{2 * time.Second, time.Sleep}
    Countdown(os.Stdout, sl)
}
