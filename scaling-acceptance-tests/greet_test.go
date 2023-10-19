package go_specs_greet_test

import (
	"testing"

	go_specs_greet "github.com/isedaniel/go-specs-greet"
	"github.com/isedaniel/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecifications(
		t, 
		specifications.GreeterAdapter(go_specs_greet.Greet))
}
