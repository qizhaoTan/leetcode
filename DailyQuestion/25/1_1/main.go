package main

import (
	"slices"
	"strconv"
	"strings"
)

func Map[S ~[]E, E, F any](s S, f func(e E) F) []F {
	if len(s) == 0 {
		return nil
	}
	ret := make([]F, 0, len(s))
	for _, e := range s {
		ret = append(ret, f(e))
	}
	return ret
}

func toBinary(i int) (ret []byte) {
	for i > 0 {
		if i&1 == 0 {
			ret = append(ret, '0')
		} else {
			ret = append(ret, '1')
		}
		i >>= 1
	}
	slices.Reverse(ret)
	return ret
}

func convertDateToBinary(date string) string {
	arr := strings.Split(date, "-")
	arr = Map(arr, func(str string) string {
		i, _ := strconv.Atoi(str)
		return string(toBinary(i))
	})
	return strings.Join(arr, "-")
}
