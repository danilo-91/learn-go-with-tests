package main

import (
	"go-with-tests/maths/clock"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clock.SVGWriter(os.Stdout, t)
}
