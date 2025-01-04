package __2

import (
	"slices"
)

type section struct {
	start int
	end   int
}

type MyCalendar struct {
	sections []section
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (s *MyCalendar) Book(startTime int, endTime int) bool {
	if slices.ContainsFunc(s.sections, func(sec section) bool {
		if sec.start <= startTime && startTime < sec.end {
			return true
		}
		if startTime <= sec.start && sec.start < endTime {
			return true
		}
		return false
	}) {
		return false
	}

	s.sections = append(s.sections, section{start: startTime, end: endTime})
	return true
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
