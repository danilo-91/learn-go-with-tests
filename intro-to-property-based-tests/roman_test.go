package roman

import (
	"fmt"
	"testing"
)

var cases = []struct {
	arabic int
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{11, "XI"},
	{12, "XII"},
	{13, "XIII"},
	{14, "XIV"},
	{15, "XV"},
	{16, "XVI"},
	{17, "XVII"},
	{18, "XVIII"},
	{19, "XIX"},
	{20, "XX"},
	{25, "XXV"},
	{30, "XXX"},
	{36, "XXXVI"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestAToRoman(t *testing.T) {
	for _, test := range cases {
		t.Run(test.roman, func(t *testing.T) {
			got := AToRoman(test.arabic)
			want := test.roman
			assertString(t, got, want)
		})
	}
}

func TestRToArabic(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%q to %d", test.roman, test.arabic), func(t *testing.T) {
			got := RToArabic(test.roman)
			want := test.arabic
			assertInt(t, got, want)
		})
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d but wanted %d", got, want)
	}
}
