package __31

func percentageLetter(s string, letter byte) int {
	dst := int32(letter)

	var a, b int
	for _, v := range s {
		b++
		if v == dst {
			a++
		}
	}

	return a * 100 / b
}
