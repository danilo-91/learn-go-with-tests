package specifications

type GreeterAdapter func(name string) string

func (g GreeterAdapter) Greet(name string) (string, error) {
	return g(name), nil
}
