package roman

import (
	"testing"
)

func TestAToRoman(t *testing.T) {
    cases := []struct {
        Description string
        n int
        roman string
    }{
        {"1 to I", 1, "I"},
        {"2 to II", 2, "II"},
        {"4 to IV", 4, "IV"},
        {"5 to V", 5, "V"},
        {"6 to VI", 6, "VI"},
        {"7 to VII", 7, "VII"},
        {"8 to VIII", 8, "VIII"},
        {"9 to IX", 9, "IX"},
        {"10 to X", 10, "X"},
        {"11 to XI", 11, "XI"},
        {"12 to XII", 12, "XII"},
        {"13 to XIII", 13, "XIII"},
        {"14 to XIV", 14, "XIV"},
        {"15 to XV", 15, "XV"},
        {"16 to XVI", 16, "XVI"},
        {"17 to XVII", 17, "XVII"},
        {"18 to XVIII", 18, "XVIII"},
        {"19 to XIX", 19, "XIX"},
        {"20 to XX", 20, "XX"},
        {"25 to XXV", 25, "XXV"},
        {"30 to XXX", 30, "XXX"},
        {"36 to XXXVI", 36, "XXXVI"},
    }

    for _, test := range cases {
        t.Run(test.Description, func(t *testing.T) {
            got := AToRoman(test.n)
            want := test.roman
            assertString(t, got, want)
        })
    }
}

func assertString(t testing.TB, got, want string) {
    t.Helper()
    if got != want {
        t.Errorf("got %q, but wanted %q", got, want)
    }
}

