package __2

func maximumTripletValue(nums []int) (ret int64) {
	i, j, k := 0, 1, 2
	for {
		ret = max(ret, int64((nums[i]-nums[j])*nums[k]))
		k++
		if k == len(nums) {
			j++
			if j == len(nums)-1 {
				i++
				if i == len(nums)-2 {
					break
				}
				j = i + 1
			}
			k = j + 1
		}
	}

	return ret
}
