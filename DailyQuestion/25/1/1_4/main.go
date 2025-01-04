package __4

type section struct {
	start int
	end   int
	cnt   int
	next  *section
}

type MyCalendarThree struct {
	max  int
	head *section
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		max:  1,
		head: &section{start: -1, end: -1, cnt: 0, next: nil},
	}
}

func (s *MyCalendarThree) Book(startTime int, endTime int) int {
	s.book(startTime, endTime)
	return s.max
}

func (s *MyCalendarThree) book(startTime, endTime int) {
	p := s.head
	for ; p.next != nil; p = p.next {
		if startTime >= p.next.end {
			continue
		}

		if startTime < p.next.start {
			if endTime > p.next.start {
				p.next = &section{start: startTime, end: p.next.start, cnt: 1, next: p.next}
				startTime = p.next.end
				continue
			}

			p.next = &section{start: startTime, end: endTime, cnt: 1, next: p.next}
			return
		} else if startTime == p.next.start {
			s.max = max(s.max, p.next.cnt+1)
			if endTime > p.next.end {
				p.next.cnt++
				startTime = p.next.end
				continue
			}

			if endTime == p.next.end {
				p.next.cnt++
				return
			} else {
				p.next.start = endTime
				p.next = &section{start: startTime, end: endTime, cnt: p.next.cnt + 1, next: p.next}
				return
			}
		} else {
			p.next = &section{start: p.next.start, end: startTime, cnt: p.next.cnt, next: p.next}
			p.next.next.start = startTime
			continue
		}
	}

	p.next = &section{start: startTime, end: endTime, cnt: 1, next: nil}
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */
