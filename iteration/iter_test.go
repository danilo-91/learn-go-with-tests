package iteration

import (
    "testing"
    "fmt"
)

func TestRepeat(t *testing.T) {
    t.Run("repeat string five times", func (t *testing.T) {
        got := Repeat("a", 10)
        want := "aaaaaaaaaa"
        assertOutput(t, got, want)
    })

}

func assertOutput(t *testing.T, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("expected %q but got %q", want, got)
    }
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 10)
    }
}

func ExampleRepeat() {
    repeated := Repeat("a", 5)
    fmt.Println(repeated)
    // Output: aaaaa
}

