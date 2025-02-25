package __25

type Allocator struct {
	mems []int
}

func Constructor(n int) Allocator {
	return Allocator{mems: make([]int, n)}
}

func (s *Allocator) Allocate(size int, mID int) int {
	var startIdx, curSize int
	for i, memId := range s.mems {
		if memId != 0 {
			curSize = 0
			continue
		}

		curSize++
		if curSize == 1 {
			startIdx = i
		}

		if curSize == size {
			for j := startIdx; j <= i; j++ {
				s.mems[j] = mID
			}
			return startIdx
		}
	}

	return -1
}

func (s *Allocator) FreeMemory(mID int) (ret int) {
	for i, mem := range s.mems {
		if mem == mID {
			s.mems[i] = 0
			ret++
		}
	}

	return ret
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.FreeMemory(mID);
 */
