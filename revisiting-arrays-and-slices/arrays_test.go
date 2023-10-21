package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum an slice of ints", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		got := Sum(nums)
		want := 15
		assertInts(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum everything but the first element of a slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected '%v' but got '%v'", want, got)
		}
	})

	t.Run("sum tails but with an empty/only head slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1}, []int{1, 2, 3}, []int{4, 5, 6})
		want := []int{0, 0, 5, 11}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected '%v' but got '%v'", want, got)
		}
	})
}

func assertInts(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("expected '%v' but got '%v'", want, got)
	}
}

func ExampleSum() {
	nums := []int{3, 4, 5, 6, 7}
	sum := Sum(nums)
	fmt.Println(sum)
	// Output: 25
}

func TestReduce(t *testing.T) {
    t.Run("reduce with multiplication", func(t *testing.T) {
        fn := func(acc, el int) int {
            return acc * el
        }

        assertEqual(t, Reduce[int]([]int{1, 2, 3}, fn, 1), 6)
    })

    t.Run("reduce slice of strings", func(t *testing.T) {
        fn := func(acc, el string) string {
            return acc + el
        }

        assertEqual(t, Reduce[string]([]string{"a", "b", "c"}, fn, "."), ".abc")
    })
}

func assertEqual[T comparable](t *testing.T, got, want T) {
    t.Helper()
    if got != want {
        t.Errorf("expected '%+v', but got '%+v'", want, got)
    }
}

func TestBadBank(t *testing.T) {
    t.Run("sum transactions", func(t *testing.T) {
        trs := []Transaction{
            {
                From: "Chris",
                To: "Kata",
                Sum: 100,
            },
            {
                From: "Adil",
                To: "Chris",
                Sum: 25,
            },
        }

        assertEqual(t, BalanceFor(trs, "Kata"), 100)
        assertEqual(t, BalanceFor(trs, "Chris"), -75)
        assertEqual(t, BalanceFor(trs, "Adil"), -25)
    })
}
