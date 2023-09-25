package roman

import (
    "strings"
)

type RomanNumeral struct {
    arabic int
    roman string
}

var RomanNumerals = []RomanNumeral {
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
