package __17

func calcMinimum(queue [][]int) (ret int) {
	indexes := make([]int, len(queue))
	valFn := func(i int) int {
		return queue[i][indexes[i]]
	}

	for {
		var minIdx, maxIdx int
		for i := 1; i < len(queue); i++ {
			if valFn(i) < valFn(minIdx) {
				minIdx = i
			}
			if valFn(i) > valFn(maxIdx) {
				maxIdx = i
			}
		}

		val := valFn(maxIdx) - valFn(minIdx) + 1
		if ret == 0 || val < ret {
			ret = val
		}

		indexes[minIdx]++
		if len(queue[minIdx]) == indexes[minIdx] {
			return ret
		}
	}
}

func minimumSubarrayLength(nums []int, k int) (ret int) {
	var (
		indexes [32][]int // 每个位对应的所有下标
		queue   [][]int
	)

	// 预处理indexes 同时判断是否存在比k大的数
	for idx, num := range nums {
		if num >= k {
			return 1
		}

		var i int
		for num > 0 {
			if num&1 == 1 {
				indexes[i] = append(indexes[i], idx)
			}
			i++
			num >>= 1
		}
	}

	ret = -1
	for i := 30; i >= 0; i-- {
		if k&(1<<i) > 0 {
			if len(indexes[i]) == 0 {
				return ret
			}
			queue = append(queue, indexes[i])
		} else if len(indexes[i]) > 0 {
			queue = append(queue, indexes[i])

			val := calcMinimum(queue)
			if ret == -1 || val < ret {
				ret = val
			}

			queue = queue[:len(queue)-1]
		}
	}

	val := calcMinimum(queue)
	if ret == -1 || val < ret {
		ret = val
	}

	return ret
}
