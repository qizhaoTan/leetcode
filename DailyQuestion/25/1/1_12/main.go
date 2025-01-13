package __12

func maxInt(args ...int) int {
	if len(args) == 1 {
		return args[0]
	} else {
		return max(args[0], maxInt(args[1:]...))
	}
}

func largestCombination(candidates []int) int {
	bits := make([]int, 32)
	for i := 0; i < len(candidates); i++ {
		num := candidates[i]
		for j := 0; num > 0; j++ {
			if num&1 == 1 {
				bits[j]++
			}
			num >>= 1
		}
	}

	return maxInt(bits...)
}
