package __11

import "math"

func pow(x, n int) int {
	return int(math.Pow(float64(x), float64(n)))
}

func generateKey(num1 int, num2 int, num3 int) (ret int) {
	for i := 4; i >= 1; i-- {
		n := pow(10, i-1)
		a := num1 / n
		b := num2 / n
		c := num3 / n

		num1 %= n
		num2 %= n
		num3 %= n

		ret += min(a, b, c) * n
	}
	return ret
}
