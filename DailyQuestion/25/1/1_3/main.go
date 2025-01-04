package __3

type section struct {
	idx      int
	start    int
	end      int
	dupIdxes []int
}

type MyCalendarTwo struct {
	idx      int
	sections []section
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (s *MyCalendarTwo) Book(startTime int, endTime int) bool {
	var dupIdxes []int
	dupSet := NewSet[int]()
	for _, sec := range s.sections {
		if (sec.start <= startTime && startTime < sec.end) || (startTime <= sec.start && sec.start < endTime) {
			dupIdxes = append(dupIdxes, sec.idx)
			if dupSet.Find(sec.idx) {
				return false
			}
			dupSet.Insert(sec.idx)

			for _, idx := range sec.dupIdxes {
				if dupSet.Find(idx) {
					return false
				}
				dupSet.Insert(idx)
			}
		}
	}

	s.idx++
	s.sections = append(s.sections, section{idx: s.idx, start: startTime, end: endTime, dupIdxes: dupIdxes})
	return true
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) ToSlice() []T {
	arr := make([]T, 0, len(s))
	for i := range s {
		arr = append(arr, i)
	}
	return arr
}

func (s Set[T]) Insert(val T) {
	s[val] = struct{}{}
}

func (s Set[T]) Inserts(args []T) {
	for _, v := range args {
		s.Insert(v)
	}
}

func (s Set[T]) Find(val T) bool {
	_, ok := s[val]
	return ok
}

func (s Set[T]) Remove(val T) {
	delete(s, val)
}

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) For(fn func(key T) bool) {
	for k := range s {
		if !fn(k) {
			break
		}
	}
}
