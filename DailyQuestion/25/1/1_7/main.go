package __7

import "strings"

func countKeyChanges(s string) (ret int) {
	s = strings.ToLower(s)
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			ret++
		}
	}
	return ret
}
