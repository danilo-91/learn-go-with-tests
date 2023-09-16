package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Address string
	Age     int
}

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
		{
			"struct with two strings",
			struct {
				Name    string
				Address string
			}{"Danilo", "Santiago"},
			[]string{"Danilo", "Santiago"},
		},
		{
			"struct with one no-string field",
			struct {
				Name string
				Age  int
			}{"Danilo", 30},
			[]string{"Danilo"},
		},
		{
			"struct with nested struct",
			Person{
				"Danilo",
				Profile{"Santiago", 32},
			},
			[]string{"Danilo", "Santiago"},
		},
		{
			"pointer to struct",
			&Person{
				"Danilo",
				Profile{"Santiago", 32},
			},
			[]string{"Danilo", "Santiago"},
		},
		{
			"slice",
			[]Profile{
				{"Santiago", 32},
				{"Vancouver", 33},
			},
			[]string{"Santiago", "Vancouver"},
		},
		{
			"array",
			[2]Profile{
				{"Santiago", 32},
				{"Vancouver", 33},
			},
			[]string{"Santiago", "Vancouver"},
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

	t.Run("map[string]string", func(t *testing.T) {
		var got []string

		m := map[string]string{
			"foo": "foo",
			"bar": "bar",
			"baz": "baz",
		}

		walk(m, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "foo")
		assertContains(t, got, "bar")
		assertContains(t, got, "baz")

	})
}

func assertContains(t *testing.T, got []string, want string) {
	t.Helper()
	for _, s := range got {
		if s == want {
			return
		}
	}
	t.Errorf("expected %q in map, but map only has %v", want, got)
}
