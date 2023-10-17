package main

import (
	go_specs_greet "github.com/isedaniel/go-specs-greet"
	"github.com/isedaniel/go-specs-greet/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecifications(t, driver)
}
