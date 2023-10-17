package main

import (
	"testing"
	"github.com/isedaniel/go-specs-greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	specifications.GreetSpecifications(t, nil)
}
