package main

import (
	"math"
	"strconv"
)

type key struct {
	n, start, finish int64
}

type dp struct {
	sn, s, limit int64

	m map[key]int64
}

// n  表示finish的位
// sn 表示s的位
func (s *dp) f(n, start, finish int64) int64 {
	if n == s.sn {
		if start <= s.s && s.s <= finish {
			return 1
		}

		return 0
	}

	k := key{n: n, start: start, finish: finish}
	if v, ok := s.m[k]; ok {
		return v
	}

	tmp := int64(math.Pow(10, float64(n-2)))
	a := start / tmp
	if a > s.limit {
		return 0
	}

	var ret int64
	b := finish / tmp
	if s.limit < b {
		ret = s.f(n-1, start%tmp, tmp-1) + (s.limit-a)*s.f(n-1, 0, tmp-1)
	} else {
		ret = s.f(n-1, start%tmp, tmp-1) + (b-a-1)*s.f(n-1, 0, tmp-1) + s.f(n-1, 0, finish%tmp)
	}

	s.m[k] = ret
	return ret
}

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	si, _ := strconv.Atoi(s)
	si6 := int64(si)
	n, sn := int64(1), int64(1)
	for i := finish; i > 0; i = i / 10 {
		n++
	}
	for i := si6; i > 0; i = i / 10 {
		sn++
	}

	v := &dp{
		sn:    sn,
		s:     si6,
		limit: int64(limit),
		m:     make(map[key]int64),
	}
	return v.f(n, start, finish)
}
