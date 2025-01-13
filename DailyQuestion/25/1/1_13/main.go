package __13

func Sum[S ~[]E, E int](s S) (ret E) {
	for _, e := range s {
		ret += e
	}
	return ret
}

func waysToSplitArray(nums []int) (result int) {
	total := Sum(nums)
	left, right := 0, total
	for i := 0; i < len(nums)-1; i++ {
		left += nums[i]
		right -= nums[i]
		if left >= right {
			result++
		}
	}
	return result
}
