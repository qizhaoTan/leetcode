package __3

func maximumTripletValue(nums []int) (ret int64) {
	var e int
	lMaxs := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		e = max(e, nums[i])
		lMaxs[i+1] = e
	}

	e = 0
	rMaxs := make([]int, len(nums))
	for i := len(nums) - 1; i > 0; i-- {
		e = max(e, nums[i])
		rMaxs[i-1] = e
	}

	for j := 1; j < len(nums)-1; j++ {
		le := lMaxs[j]
		re := rMaxs[j]

		ret = max(ret, int64(le-nums[j])*int64(re))
	}

	return ret
}
