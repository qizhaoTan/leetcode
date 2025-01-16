package __16

func minimumSubarrayLength(nums []int, k int) int {
	n := len(nums)
	for i := 1; i <= n; i++ {
		for j := 0; j < n-i+1; j++ {
			var val int
			for l := 0; l < i; l++ {
				val |= nums[l+j]
			}
			if val >= k {
				return i
			}
		}
	}

	return -1
}
