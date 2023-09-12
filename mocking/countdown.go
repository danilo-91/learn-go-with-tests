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

type DefaultSleeper struct {}

func (sl *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, sl Sleeper) {
    for i := countdownStart; i > 0; i-- {
        fmt.Fprintln(w, i)
        sl.Sleep()
    }
    fmt.Fprintln(w, lastLine)
}

func main() {
    sl := &DefaultSleeper{}
    Countdown(os.Stdout, sl)
}
