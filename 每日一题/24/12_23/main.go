package main

import (
	"fmt"
	"math/rand"
)

type Segment struct {
	start, end int
	free       bool

	prev *Segment
	next *Segment
}

func newSegment(start, end int, free bool) *Segment {
	if start > end {
		return nil
	}

	return &Segment{start: start, end: end, free: free}
}

func (s *Segment) Dist() int {
	p := (s.start + s.end) >> 1
	dist := p - s.start - 1
	return dist
}

func (s *Segment) Size() int {
	return s.end - s.start
}

func cmp(a, b int) int {
	if a == b {
		return 0
	}

	if a < b {
		return 1
	}

	return -1
}

func cmpSegment(a, b *Segment) int {
	distA := a.Dist()
	distB := b.Dist()
	if distA != distB {
		return cmp(distA, distB)
	}

	return cmp(b.start, a.start)
}

type ExamRoom struct {
	n    int
	m    map[int]*Segment
	sl   *SkipList[*Segment]
	head *Segment
	tail *Segment
}

func Constructor(n int) ExamRoom {
	segment := &Segment{start: 0, end: n - 1, free: true, prev: nil, next: nil}
	ret := ExamRoom{
		n:    n,
		m:    make(map[int]*Segment),
		sl:   NewSkipList[*Segment](cmpSegment),
		head: segment,
		tail: segment,
	}
	ret.sl.Insert(segment)
	return ret
}

func (s *ExamRoom) Seat() int {
	segment := s.sl.Front().Value

	p := (segment.start + segment.end) >> 1
	dist := max(p-segment.start-1, 0)

	if s.head.free && s.head.Size() > dist {
		segment = s.head
		p = s.head.start
		dist = s.head.Size()
	}

	if s.tail.free && s.tail.Size() > dist {
		segment = s.tail
		p = s.tail.end
	}

	s.seat(segment, p)

	return p
}

func (s *ExamRoom) add(segment *Segment) {
	if segment == nil {
		return
	}
	s.sl.Insert(segment)
}

func (s *ExamRoom) del(segment *Segment) {
	if segment == nil {
		return
	}
	if s.head == segment {
		s.head = nil
	}
	if s.tail == segment {
		s.tail = nil
	}
	if !segment.free {
		return
	}

	s.sl.Delete(segment)
}

func (s *ExamRoom) splice(a, b *Segment) *Segment {
	if a == nil && b == nil {
		return nil
	}

	if a == nil {
		b.prev = nil
		return b
	}

	if b == nil {
		a.next = nil
		return a
	}

	a.next = b
	b.prev = a
	return b
}

func (s *ExamRoom) seat(segment *Segment, p int) {
	s.del(segment)

	pprev := segment.prev
	prev := newSegment(segment.start, p-1, true)
	pSegment := newSegment(p, p, false)
	next := newSegment(p+1, segment.end, true)
	nnext := segment.next

	tmp := s.splice(pprev, prev)
	tmp = s.splice(tmp, pSegment)
	tmp = s.splice(tmp, next)
	tmp = s.splice(tmp, nnext)

	s.m[p] = pSegment
	s.add(prev)
	s.add(next)

	s.tryResume(prev, pSegment, next)
}

func (s *ExamRoom) Leave(p int) {
	segment := s.m[p]
	delete(s.m, p)
	s.del(segment)

	prev := segment.prev
	next := segment.next

	s.del(prev)
	s.del(next)

	start := segment.start
	end := segment.end
	if prev != nil && prev.free {
		start = prev.start
		prev = prev.prev
	}
	if next != nil && next.free {
		end = next.end
		next = next.next
	}
	nSegment := newSegment(start, end, true)
	tmp := s.splice(prev, nSegment)
	tmp = s.splice(tmp, next)
	s.add(nSegment)

	s.tryResume(prev, nSegment, next)
}

func (s *ExamRoom) tryResume(prev, segment, next *Segment) {
	if s.head == nil {
		if prev == nil {
			s.head = segment
		} else {
			s.head = prev
		}
	}
	if s.tail == nil {
		if next == nil {
			s.tail = segment
		} else {
			s.tail = next
		}
	}
}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */

type SkipList[T any] struct {
	header *Element[T]
	tail   *Element[T]
	update []*Element[T]
	rank   []int
	length int
	level  int

	compare func(T, T) int
}

// New returns an initialized skiplist.
func NewSkipList[T any](compare func(T, T) int) *SkipList[T] {
	return &SkipList[T]{
		header:  createElement[T](SKIPLIST_MAXLEVEL),
		tail:    nil,
		update:  make([]*Element[T], SKIPLIST_MAXLEVEL),
		rank:    make([]int, SKIPLIST_MAXLEVEL),
		length:  0,
		level:   1,
		compare: compare,
	}
}

// Init initializes or clears skiplist sl.
func (s *SkipList[T]) Clear() *SkipList[T] {
	s.header = createElement[T](SKIPLIST_MAXLEVEL)
	s.tail = nil
	s.update = make([]*Element[T], SKIPLIST_MAXLEVEL)
	s.rank = make([]int, SKIPLIST_MAXLEVEL)
	s.length = 0
	s.level = 1
	return s
}

// Front returns the first elements of skiplist sl or nil.
func (s *SkipList[T]) Front() *Element[T] {
	return s.header.level[0].forward
}

// Back returns the last elements of skiplist sl or nil.
func (s *SkipList[T]) Back() *Element[T] {
	return s.tail
}

// Len returns the numbler of elements of skiplist sl.
func (s *SkipList[T]) Len() int {
	return s.length
}

// Insert inserts v, increments sl.length, and returns a new element of wrap v.
func (s *SkipList[T]) Insert(v T) *Element[T] {
	s.Delete(v)
	x := s.header
	for i := s.level - 1; i >= 0; i-- {
		// store rank that is crossed to reach the insert position
		if i == s.level-1 {
			s.rank[i] = 0
		} else {
			s.rank[i] = s.rank[i+1]
		}
		//for x.level[i].forward != nil && x.level[i].forward.Value.Less(v) {
		for x.level[i].forward != nil && s.compare(x.level[i].forward.Value, v) == -1 {
			s.rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		s.update[i] = x
	}

	// ensure that the v is unique, the re-insertion of v should never happen since the
	// caller of sl.Insert() should test in the hash table if the element is already inside or not.
	level := randomLevel()
	if level > s.level {
		for i := s.level; i < level; i++ {
			s.rank[i] = 0
			s.update[i] = s.header
			s.update[i].level[i].span = s.length
		}
		s.level = level
	}

	x = newElement(level, v)
	for i := 0; i < level; i++ {
		x.level[i].forward = s.update[i].level[i].forward
		s.update[i].level[i].forward = x

		// update span covered by update[i] as x is inserted here
		x.level[i].span = s.update[i].level[i].span - s.rank[0] + s.rank[i]
		s.update[i].level[i].span = s.rank[0] - s.rank[i] + 1
	}

	// increment span for untouched levels
	for i := level; i < s.level; i++ {
		s.update[i].level[i].span++
	}

	if s.update[0] == s.header {
		x.backward = nil
	} else {
		x.backward = s.update[0]
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		s.tail = x
	}
	s.length++

	return x
}

// deleteElement deletes e from its skiplist, and decrements sl.length.
func (s *SkipList[T]) deleteElement(e *Element[T], update []*Element[T]) {
	for i := 0; i < s.level; i++ {
		if update[i].level[i].forward == e {
			update[i].level[i].span += e.level[i].span - 1
			update[i].level[i].forward = e.level[i].forward
		} else {
			update[i].level[i].span -= 1
		}
	}

	if e.level[0].forward != nil {
		e.level[0].forward.backward = e.backward
	} else {
		s.tail = e.backward
	}

	for s.level > 1 && s.header.level[s.level-1].forward == nil {
		s.level--
	}
	s.length--
}

// Remove removes e from sl if e is an element of skiplist sl.
// It returns the element value e.Value.
func (s *SkipList[T]) Remove(e *Element[T]) bool {
	x := s.find(e.Value) // x.Value >= e.Value
	//if x == e && !e.Value.Less(x.Value) { // e.Value >= x.Value
	if x == e && s.compare(e.Value, x.Value) >= 0 { // e.Value >= x.Value
		s.deleteElement(x, s.update)
		return true
	}

	return false
}

// Delete deletes an element e that e.Value == v, and returns e.Value or nil.
func (s *SkipList[T]) Delete(v T) bool {
	x := s.find(v) // x.Value >= v
	//if x != nil && !v.Less(x.Value) { // v >= x.Value
	if x != nil && s.compare(v, x.Value) >= 0 { // v >= x.Value
		s.deleteElement(x, s.update)
		return true
	}

	return false
}

// FindFunc 找到第一个元素e，使得less(e.Value) == false
func (s *SkipList[T]) FindFunc(less func(v T) bool) *Element[T] {
	return s.findFunc(less)
}

// findFunc  找到第一个元素e，使得less(e.Value) == false
func (s *SkipList[T]) findFunc(less func(v T) bool) *Element[T] {
	x := s.header
	for i := s.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && less(x.level[i].forward.Value) {
			x = x.level[i].forward
		}
		s.update[i] = x
	}

	return x.level[0].forward
}

// Find finds an element e that e.Value == v, and returns e or nil.
func (s *SkipList[T]) Find(v T) *Element[T] {
	x := s.find(v) // x.Value >= v
	//if x != nil && !v.Less(x.Value) { // v >= x.Value
	if x != nil && s.compare(v, x.Value) >= 0 { // v >= x.Value
		return x
	}

	return nil
}

// find finds the first element e that e.Value >= v, and returns e or nil.
func (s *SkipList[T]) find(v T) *Element[T] {
	x := s.header
	for i := s.level - 1; i >= 0; i-- {
		//for x.level[i].forward != nil && x.level[i].forward.Value.Less(v) {
		for x.level[i].forward != nil && s.compare(x.level[i].forward.Value, v) == -1 {
			x = x.level[i].forward
		}
		s.update[i] = x
	}

	return x.level[0].forward
}

// GetRank finds the rank for an element e that e.Value == v,
// Returns 0 when the element cannot be found, rank otherwise.
// Note that the rank is 1-based due to the span of sl.header to the first element.
func (s *SkipList[T]) GetRank(v T) int {
	x := s.header
	rank := 0
	for i := s.level - 1; i >= 0; i-- {
		//for x.level[i].forward != nil && x.level[i].forward.Value.Less(v) {
		for x.level[i].forward != nil && s.compare(x.level[i].forward.Value, v) == -1 {
			rank += x.level[i].span
			x = x.level[i].forward
		}
		//if x.level[i].forward != nil && !x.level[i].forward.Value.Less(v) && !v.Less(x.level[i].forward.Value) {
		if x.level[i].forward != nil && s.compare(x.level[i].forward.Value, v) == 0 {
			rank += x.level[i].span
			return rank
		}
	}

	return 0
}

// GetElementByRank finds an element by ites rank. The rank argument needs bo be 1-based.
// Note that is the first element e that GetRank(e.Value) == rank, and returns e or nil.
func (s *SkipList[T]) GetElementByRank(rank int) *Element[T] {
	x := s.header
	traversed := 0
	for i := s.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && traversed+x.level[i].span <= rank {
			traversed += x.level[i].span
			x = x.level[i].forward
		}
		if traversed == rank {
			return x
		}
	}

	return nil
}

// SubList returns a sub list of sl, which elements' value are in [i, j].
func (s *SkipList[T]) SubList(i, j int) []T {
	if i <= 0 {
		i = 1
	}

	ret := make([]T, 0, j-i+1)
	e := s.GetElementByRank(i)
	for ; i <= j; i++ {
		if e == nil {
			break
		}
		ret = append(ret, e.Value)
		e = e.Next()
	}
	return ret
}

func (s *SkipList[T]) AllList() []T {
	return s.SubList(1, s.Len())
}

const SKIPLIST_MAXLEVEL = 32
const SKIPLIST_BRANCH = 4

type skiplistLevel[T any] struct {
	forward *Element[T]
	span    int
}

type Element[T any] struct {
	Value    T
	backward *Element[T]
	level    []*skiplistLevel[T]
}

// Next returns the next skiplist element or nil.
func (e *Element[T]) Next() *Element[T] {
	return e.level[0].forward
}

// Prev returns the previous skiplist element of nil.
func (e *Element[T]) Prev() *Element[T] {
	return e.backward
}

// newElement returns an initialized element.
func newElement[T any](level int, v T) *Element[T] {
	slLevels := make([]*skiplistLevel[T], level)
	for i := 0; i < level; i++ {
		slLevels[i] = new(skiplistLevel[T])
	}

	return &Element[T]{
		Value:    v,
		backward: nil,
		level:    slLevels,
	}
}

func createElement[T any](level int) *Element[T] {
	slLevels := make([]*skiplistLevel[T], level)
	for i := 0; i < level; i++ {
		slLevels[i] = new(skiplistLevel[T])
	}

	return &Element[T]{
		backward: nil,
		level:    slLevels,
	}
}

// randomLevel returns a random level.
func randomLevel() int {
	level := 1
	for (rand.Int31()&0xFFFF)%SKIPLIST_BRANCH == 0 {
		level += 1
	}

	if level < SKIPLIST_MAXLEVEL {
		return level
	} else {
		return SKIPLIST_MAXLEVEL
	}
}

func main() {
	s := Constructor(4)
	println(s.Seat())
	test(s)
	println(s.Seat())
	test(s)
	println(s.Seat())
	test(s)
	println(s.Seat())
	test(s)
	s.Leave(1)
	test(s)
	s.Leave(3)
	test(s)
	println(s.Seat())
	test(s)
}

func test(s ExamRoom) {
	arr := s.sl.AllList()
	fmt.Println(s.head)
	fmt.Println(s.tail)
	for _, segment := range arr {
		fmt.Println(segment)
	}
	println("-----")
}
