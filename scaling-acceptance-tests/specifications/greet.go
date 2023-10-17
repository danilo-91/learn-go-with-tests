package specifications

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

type Greeter interface {
	Greet() (string, error)
}

func GreetSpecifications(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet()
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, World")
}
