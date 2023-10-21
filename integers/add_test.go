package integers

import (
    "testing"
    "fmt"
)

func TestAdd(t *testing.T) {
    t.Run("2 + 2 = 4", func (t *testing.T) {
        got := Add(2, 2)
        want := 4

        if got != want {
            t.Errorf("got '%d' want '%d'", got, want)
        }
    })
}

func ExampleAdd() {
    sum := Add(17, 2)
    fmt.Println(sum)
    // Output: 19
}
