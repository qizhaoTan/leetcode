package __19

import "errors"

type Item struct {
	idx int
	min int
	max int
}

func extremum(arr []int) (int, int) {
	minVal, maxVal := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if minVal > arr[i] {
			minVal = arr[i]
		}

		if maxVal < arr[i] {
			maxVal = arr[i]
		}
	}
	return minVal, maxVal
}

func maxDistance(arrays [][]int) int {
	minHeap, _ := NewAny(func(a, b *Item) bool {
		return a.min < b.min
	})
	maxHeap, _ := NewAny(func(a, b *Item) bool {
		return a.max > b.max
	})

	for idx, arr := range arrays {
		minVal, maxVal := extremum(arr)
		item := &Item{idx: idx, min: minVal, max: maxVal}
		minHeap.Push(item)
		maxHeap.Push(item)
	}

	minItem, maxItem := minHeap.Top(), maxHeap.Top()
	if minItem.idx != maxItem.idx {
		return maxItem.max - minItem.min
	}

	minHeap.Pop()
	maxHeap.Pop()

	minItem2, maxItem2 := minHeap.Top(), maxHeap.Top()
	ret1 := maxItem.max - minItem2.min
	ret2 := maxItem2.max - minItem.min
	if ret1 > ret2 {
		return ret1
	}

	return ret2
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
