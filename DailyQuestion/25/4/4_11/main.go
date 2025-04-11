package __11

func isSymmetric(v int) bool {
	arr := make([]int, 0, 4)
	for v > 0 {
		i := v % 10
		arr = append(arr, i)
		v = v / 10
	}
	if len(arr)%2 != 0 {
		return false
	}

	var a, b int
	for i := 0; i < len(arr)/2; i++ {
		a += arr[i]
		b += arr[len(arr)-i-1]
	}
	return a == b
}

func countSymmetricIntegers(low int, high int) (ret int) {
	for i := low; i <= high; i++ {
		if isSymmetric(i) {
			ret++
		}
	}
	return ret
}
