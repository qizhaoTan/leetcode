package __8

import (
	"strconv"
	"strings"
)

func largestGoodInteger(num string) string {
	for i := 9; i >= 0; i-- {
		str := strings.Repeat(strconv.Itoa(i), 3)
		if strings.Contains(num, str) {
			return str
		}
	}
	return ""
}
