package roman

import (
	"strings"
)

type RomanNumeral struct {
	arabic int
	roman  string
}

var RomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func AToRoman(arabic int) string {
	var roman strings.Builder
	for _, numeral := range RomanNumerals {
		for arabic >= numeral.arabic {
			roman.WriteString(numeral.roman)
			arabic -= numeral.arabic
		}
	}
	return roman.String()
}

func RToArabic(roman string) int {
	arabic := 1
	for r := AToRoman(arabic); !(roman == r); r = AToRoman(arabic) {
		arabic++
	}
	return arabic
}
