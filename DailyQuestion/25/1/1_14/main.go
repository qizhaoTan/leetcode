package __14

func minOperations(nums []int, k int) (result int) {
	for _, num := range nums {
		if num < k {
			result++
		}
	}
	return result
}
