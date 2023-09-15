package main

import (
	"testing"
    "reflect"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		name     string
		input    interface{}
		expected []string
	}{
		{
			"struct with one string",
			struct {
				Name string
			}{"Danilo"},
			[]string{"Danilo"},
		},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			var got []string

			walk(test.input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.expected) {
				t.Errorf("expected %v, but got %v", test.expected, got)
			}
		})
	}
}
