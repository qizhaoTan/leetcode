package _8_RK

import "strings"

func hash(str string) (ret int) {
	for _, c := range str {
		ret += int(c - 'a')
	}
	return ret
}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}

	n := len(needle)
	needleHash := hash(needle)
	hayHash := hash(haystack[:n])

	if hayHash == needleHash && strings.Compare(haystack[:n], needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack)-n; i++ {
		hayHash -= int(haystack[i] - 'a')
		hayHash += int(haystack[i+n] - 'a')
		if hayHash == needleHash && strings.Compare(haystack[i+1:i+n+1], needle) == 0 {
			return i + 1
		}
	}

	return -1
}
