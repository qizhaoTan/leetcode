package __6

import "slices"

func maxConsecutive(bottom int, top int, special []int) int {
	slices.Sort(special)
	ret := special[0] - bottom
	for i := 1; i < len(special); i++ {
		ret = max(ret, special[i]-special[i-1]-1)
	}
	ret = max(ret, top-special[len(special)-1])
	return ret
}
