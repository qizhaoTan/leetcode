package __6

import "slices"

func largestDivisibleSubset(nums []int) []int {
	numsClone := slices.Clone(nums)
	slices.Sort(numsClone)

	for i := 0; i < len(nums); i++ {
		var tmp []int
		for j := i + 1; j < len(nums); j++ {
			//if nums[j] % ni
		}
	}
}
