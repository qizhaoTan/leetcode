package __27

type Node struct {
	str []byte

	pre  *Node
	next *Node
}

type TextEditor struct {
	idx int
	cur *Node
}

func Constructor() TextEditor {
	return TextEditor{cur: &Node{}, idx: 0}
}

func (s *TextEditor) AddText(text string) {
	s.split()

	s.cur.str = append(s.cur.str, []byte(text)...)
	s.idx += len(text)
}

func (s *TextEditor) DeleteText(k int) (ret int) {
	s.split()

	for k > s.idx {
		ret += s.idx
		k -= s.idx
		s.cur.str = nil
		s.idx = 0

		if s.cur.pre == nil {
			break
		}

		s.toPre()
	}

	if k > s.idx {
		k = s.idx
	}

	ret += k
	s.idx -= k
	s.cur.str = s.cur.str[:s.idx:s.idx]
	return ret
}

func (s *TextEditor) CursorLeft(k int) string {
	for k > s.idx {
		k -= s.idx
		s.idx = 0

		if s.cur.pre == nil {
			break
		}

		s.toPre()
	}

	if k > s.idx {
		k = s.idx
	}

	s.idx -= k
	return s.ret()
}

func (s *TextEditor) CursorRight(k int) string {
	for k > len(s.cur.str)-s.idx {
		k -= len(s.cur.str) - s.idx
		s.idx = len(s.cur.str)

		if s.cur.next == nil {
			break
		}

		s.toNext()
	}

	if k > len(s.cur.str)-s.idx {
		k = len(s.cur.str) - s.idx
	}

	s.idx += k
	return s.ret()
}

// 把当前下标后面的内容截断
func (s *TextEditor) split() {
	if len(s.cur.str) != s.idx {
		str := s.cur.str[s.idx:len(s.cur.str):len(s.cur.str)]
		s.cur.next = &Node{str: str, next: s.cur.next, pre: s.cur}
		s.cur.str = s.cur.str[:s.idx:s.idx]
	}
}

func (s *TextEditor) ret() string {
	if s.idx >= 10 {
		return string(s.cur.str[s.idx-10 : s.idx])
	}

	return ret(s.cur.pre, 10-s.idx) + string(s.cur.str[:s.idx])
}

func ret(cur *Node, n int) string {
	if cur == nil {
		return ""
	}

	if len(cur.str) >= n {
		return string(cur.str[len(cur.str)-n:])
	}

	return ret(cur.pre, n-len(cur.str)) + string(cur.str)
}

func (s *TextEditor) toPre() {
	if s.idx != 0 {
		panic("toNext")
	}

	s.tryDelCur()
	s.cur = s.cur.pre
	s.idx = len(s.cur.str)
}

func (s *TextEditor) toNext() {
	if len(s.cur.str) != s.idx {
		panic("toNext")
	}

	s.tryDelCur()
	s.cur = s.cur.next
	s.idx = 0
}

func (s *TextEditor) tryDelCur() {
	if len(s.cur.str) != 0 {
		return
	}

	if s.cur.pre != nil {
		s.cur.pre.next = s.cur.next
	}

	if s.cur.next != nil {
		s.cur.next.pre = s.cur.pre
	}
}

/**
 * Your TextEditor object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddText(text);
 * param_2 := obj.DeleteText(k);
 * param_3 := obj.CursorLeft(k);
 * param_4 := obj.CursorRight(k);
 */
