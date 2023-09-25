package roman

func NToRoman(n int) string {
	roman := ""
	for n > 0 {
		roman += "I"
		n -= 1
	}
	return roman

}
