package __8

func minimumOperations(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	var cnt int
	for len(m) != len(nums) {
		n := 3
		if len(nums) < n {
			n = len(nums)
		}

		for _, v := range nums[:n] {
			m[v]--
			if m[v] == 0 {
				delete(m, v)
			}
		}

		nums = nums[n:]
		cnt++
	}

	return cnt
}
