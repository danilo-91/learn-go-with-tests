package generics_test

import (
	"testing"
	"github.com/isedaniel/generics"
)

func TestStack(t *testing.T) {
    t.Run("integer stack", func(t *testing.T) {
        s := new(generics.Stack[int])        

        // check stack empty
        AssertTrue(t, s.IsEmpty())

        // add int, check not empty
        s.Push(1)
        AssertFalse(t, s.IsEmpty())

        // add another, pop back
        s.Push(2)
        v, _ := s.Pop()
        AssertEqual(t, v, 2)
        v, _ = s.Pop()
        AssertEqual(t, v, 1)
        AssertTrue(t, s.IsEmpty())

        // pop returns int instead of interface{}
        s.Push(1)
        s.Push(2)
        first, _ := s.Pop()
        second, _ := s.Pop()
        AssertEqual(t, first + second, 3)
    })
}

func AssertTrue(t *testing.T, got bool) {
    t.Helper()
    if !got {
        t.Errorf("got %v, want true", got)
    }
}

func AssertFalse(t *testing.T, got bool) {
    t.Helper()
    if got {
        t.Errorf("got %v, wanted false", got)
    }
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
    t.Helper()
    if got != want {
        t.Errorf("got %+v, but expected %+v", got, want)
    }
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
    t.Helper()
    if got == want {
        t.Errorf("didn't want '%+v'", want)
    }
}
