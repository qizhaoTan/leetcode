package __7

import "slices"

func canPartition(nums []int) bool {
	var total int
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	dst := total / 2
	m := make(map[int]struct{})
	m[0] = struct{}{}

	slices.Sort(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		tmp := make([]int, 0, len(m))
		for val := range m {
			if val > dst {
				continue
			}
			if val == dst {
				return true
			}
			tmp = append(tmp, val+nums[i])
		}

		for _, val := range tmp {
			m[val] = struct{}{}
		}
	}

	return false
}
