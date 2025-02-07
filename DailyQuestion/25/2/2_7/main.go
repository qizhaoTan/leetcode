package __7

func generateMatrix(n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}

	num := 0
	for i := 0; i < n; i++ {
		if i == n-i-1 {
			num++
			ret[i][i] = num
			break
		}

		for x, y := i, i; x < n-i-1; x++ {
			num++
			ret[y][x] = num
		}

		for x, y := n-i-1, i; y < n-i-1; y++ {
			num++
			ret[y][x] = num
		}

		for x, y := n-i-1, n-i-1; x > i; x-- {
			num++
			ret[y][x] = num
		}

		for x, y := i, n-i-1; y > i; y-- {
			num++
			ret[y][x] = num
		}
	}

	return ret
}
