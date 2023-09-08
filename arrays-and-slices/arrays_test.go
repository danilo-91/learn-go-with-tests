package main

import (
    "testing"
    "fmt"
    "reflect"
)

func TestSum(t *testing.T) {
    t.Run("sum an slice of ints", func (t *testing.T) {
        nums := []int{1, 2, 3, 4, 5}
        got := Sum(nums)
        want := 15
        assertInts(t, got, want)
    })
}

func TestSumAll(t *testing.T) {
    t.Run("sum each of an slice of int slices", func (t *testing.T) {
        got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
        want := []int{6, 15}
        if !reflect.DeepEqual(got, want) {
            t.Errorf("expected '%v' but got '%v'", want, got)
        }
    })
}

func TestSumAllTails(t *testing.T) {
    t.Run("sum everything but the first element of a slice", func (t *testing.T) {
        got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
        want := []int{5, 11}
        if !reflect.DeepEqual(got, want) {
            t.Errorf("expected '%v' but got '%v'", want, got)
        }
    })

    t.Run("sum tails but with an empty/only head slice", func (t *testing.T) {
        got := SumAllTails([]int{}, []int{1}, []int{1, 2, 3}, []int{4, 5, 6})
        want := []int{0, 0, 5, 11}
        if !reflect.DeepEqual(got ,want) {
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

