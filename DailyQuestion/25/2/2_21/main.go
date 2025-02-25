package __21

func minimumWhiteTiles(floor string, numCarpets int, carpetLen int) int {
	if carpetLen >= len(floor) {
		return 0
	}

	n := 0
	for i := 0; i < carpetLen; i++ {
		n += int(floor[i] - '0')
	}

	arr := make([]int, 0, len(floor)-carpetLen)
	arr[0] = n
	for i := 1; i < len(floor)-carpetLen; i++ {
		n -= int(floor[i-1] - '0')
		n += int(floor[i+carpetLen-1] - '0')
		arr[i] = n
	}

	var idx, maxCnt int
	for i, clearCnt := range arr {
		if clearCnt > maxCnt {
			idx = i
			maxCnt = clearCnt
		}
	}

	if maxCnt == 0 {
		return 0
	}

	arr[idx] = 0
	resetNum := carpetLen - 1
	for i := idx + 1; i < len(arr) &&; i++ {
		arr[i] = 0
		resetNum--
	}
}
