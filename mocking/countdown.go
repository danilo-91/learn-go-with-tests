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

func Countdown(w io.Writer) {
    for i := countdownStart; i > 0; i-- {
        fmt.Fprintln(w, i)
        time.Sleep(1 * time.Second)
    }
    fmt.Fprintln(w, lastLine)
}

func main() {
    Countdown(os.Stdout)
}
