package _2_24

import "errors"

type Apples struct {
	day int
	cnt int
}

func eatenApples(apples []int, days []int) int {
	appleHeap, _ := NewAny[*Apples](func(a, b *Apples) bool {
		return a.day < b.day
	})

	var i, ret int
	for i = 0; i < len(apples); i++ {
		a := &Apples{day: i + days[i] - 1, cnt: apples[i]}
		if a.cnt > 0 {
			appleHeap.Push(a)
		}

		for appleHeap.Size() > 0 {
			top := appleHeap.Top()
			if top.day < i {
				appleHeap.Pop()
				continue
			}

			top.cnt--
			ret++
			if top.cnt <= 0 {
				appleHeap.Pop()
			}
			break
		}
	}

	for appleHeap.Size() > 0 {
		for appleHeap.Size() > 0 {
			top := appleHeap.Top()
			if top.day < i {
				appleHeap.Pop()
				continue
			}

			top.cnt--
			ret++
			if top.cnt <= 0 {
				appleHeap.Pop()
			}
			break
		}
		i++
	}

	return ret
}

// Heap heap implementation using generic.
type Heap[T any] struct {
	heaps    []T
	lessFunc func(a, b T) bool
}

// NewAny gives a new heap object. element can be anything, but must provide less function.
func NewAny[T any](less func(a, b T) bool) (*Heap[T], error) {
	if less == nil {
		return nil, errors.New("less func is necessary")
	}
	return &Heap[T]{
		lessFunc: less,
	}, nil
}

// Push pushes the element t onto the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Push(t T) {
	h.heaps = append(h.heaps, t)
	h.up(len(h.heaps) - 1)
}

// Top returns the minimum element (according to Less) from the heap.
// Top panics if the heap is empty.
func (h *Heap[T]) Top() T {
	return h.heaps[0]
}

// Pop removes the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
func (h *Heap[T]) Pop() {
	if len(h.heaps) <= 1 {
		h.heaps = nil
		return
	}
	h.swap(0, len(h.heaps)-1)
	h.heaps = h.heaps[:len(h.heaps)-1]
	h.down(0)
}

// Empty returns the heap is empty or not.
func (h *Heap[T]) Empty() bool {
	return len(h.heaps) == 0
}

// Size returns the size of the heap
func (h *Heap[T]) Size() int {
	return len(h.heaps)
}

func (h *Heap[T]) swap(i, j int) {
	h.heaps[i], h.heaps[j] = h.heaps[j], h.heaps[i]
}

func (h *Heap[T]) up(child int) {
	if child <= 0 {
		return
	}
	parent := (child - 1) >> 1
	if !h.lessFunc(h.heaps[child], h.heaps[parent]) {
		return
	}
	h.swap(child, parent)
	h.up(parent)
}

func (h *Heap[T]) down(parent int) {
	lessIdx := parent
	lChild, rChild := (parent<<1)+1, (parent<<1)+2
	if lChild < len(h.heaps) && h.lessFunc(h.heaps[lChild], h.heaps[lessIdx]) {
		lessIdx = lChild
	}
	if rChild < len(h.heaps) && h.lessFunc(h.heaps[rChild], h.heaps[lessIdx]) {
		lessIdx = rChild
	}
	if lessIdx == parent {
		return
	}
	h.swap(lessIdx, parent)
	h.down(lessIdx)
}
