package _8_BM

/*
第一版，实现从后往前暴力匹配

func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)
	if m < n {
		return -1
	}

	var i int
	for i <= m-n {
		j := 1
		for ; j <= n; j++ {
			if needle[n-j] != haystack[i+n-j] {
				break
			}
		}

		// 找到了
		if j == n+1 {
			return i
		}

		i++
	}
	return -1
}
*/

/*
第二版，实现加入了坏字符的规则

func calcBadChar(str string) map[uint8]int {
	ret := make(map[uint8]int)
	for i := range str {
		k := str[i] - 'a'
		ret[k] = i
	}
	return ret
}

func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)
	if m < n {
		return -1
	}

	badChar := calcBadChar(needle)

	var i int
	for i <= m-n {
		j := 1
		for ; j <= n; j++ {
			if needle[n-j] != haystack[i+n-j] {
				break
			}
		}

		// 找到了
		if j == n+1 {
			return i
		}

		k := haystack[i+n-j] - 'a'
		if idx, ok := badChar[k]; ok {
			i = i + max(1, n-j-idx)
		} else {
			i = i + n - j + 1
		}
	}
	return -1
}
*/

func calcBadChar(str string) map[uint8]int {
	ret := make(map[uint8]int)
	for i := range str {
		k := str[i] - 'a'
		ret[k] = i
	}
	return ret
}

func strStr(haystack string, needle string) int {
	m := len(haystack)
	n := len(needle)
	if m < n {
		return -1
	}

	badChar := calcBadChar(needle)

	var i int
	for i <= m-n {
		j := 1
		for ; j <= n; j++ {
			if needle[n-j] != haystack[i+n-j] {
				break
			}
		}

		// 找到了
		if j == n+1 {
			return i
		}

		k := haystack[i+n-j] - 'a'
		if idx, ok := badChar[k]; ok {
			i = i + max(1, n-j-idx)
		} else {
			i = i + n - j + 1
		}
	}
	return -1
}
