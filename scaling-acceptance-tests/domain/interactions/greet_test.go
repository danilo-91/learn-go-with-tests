package interactions_test

import (
	"testing"

	"github.com/isedaniel/go-specs-greet/domain/interactions"
	"github.com/isedaniel/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecifications(
		t, 
		specifications.GreeterAdapter(interactions.Greet))
}
