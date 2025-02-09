package __9

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	var i int
	a, b := nums[0], nums[1]
	for j := 2; j < len(nums); j++ {
		if nums[j] == nums[j-2] {
			continue
		}
		nums[i] = nums[j]
		i++
	}
	for j := i + 1; j >= 2; j-- {
		nums[j] = nums[j-2]
	}
	nums[0], nums[1] = a, b
	return i + 2
}
