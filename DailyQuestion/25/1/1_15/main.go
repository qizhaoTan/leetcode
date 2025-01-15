package __15

import "errors"

func minOperations(nums []int, k int) (result int) {
	heap, _ := NewAny(func(a, b int) bool {
		return a < b
	})

	for _, num := range nums {
		heap.Push(num)
	}

	for heap.Size() >= 2 {
		a := heap.Top()
		heap.Pop()
		b := heap.Top()
		heap.Pop()

		if a >= k {
			break
		}

		result++
		c := min(a, b)*2 + max(a, b)
		heap.Push(c)
	}

	return result
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
