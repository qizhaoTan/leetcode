package __9

func minOperations(nums []int, k int) int {
	m := make(map[int]int)
	for _, num := range nums {
		if num < k {
			return -1
		}
		m[num]++
	}

	if _, ok := m[k]; ok {
		return len(m) - 1
	}

	return len(m)
}
