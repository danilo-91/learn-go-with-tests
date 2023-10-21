package specifications

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecifications(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Danilob")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, Danilob")
}
