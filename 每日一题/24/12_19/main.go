package _2_19

func stableMountains(height []int, threshold int) (ret []int) {
	for i, v := range height[:len(height)-1] {
		if v <= threshold {
			continue
		}
		ret = append(ret, i+1)
	}
	return ret
}
