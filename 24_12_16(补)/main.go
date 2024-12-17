package main

import (
	"slices"
)

type Room struct {
	id   int
	size int

	idx  int
	prev *Room
	next *Room
	real *Room
}

func (s *Room) Real() *Room {
	if s.real == nil {
		return s
	}

	s.real = s.real.Real()
	return s.real
}

func (s *Room) RealId() int {
	if s.real == nil {
		return s.id
	}

	return s.Real().id
}

type Queries struct {
	idx  int
	id   int
	size int
}

func binarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) int {
	n := len(x)
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if cmp(x[h], target) < 0 {
			i = h + 1
		} else {
			j = h
		}
	}
	return i
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func closestRoom(rooms [][]int, queries [][]int) []int {
	n := len(rooms)
	roomIdList := make([]*Room, n)
	roomSizeList := make([]*Room, n)
	for i, val := range rooms {
		room := &Room{id: val[0], size: val[1]}
		roomIdList[i] = room
		roomSizeList[i] = room
	}

	slices.SortFunc(roomIdList, func(a, b *Room) int {
		return (*a).id - (*b).id
	})
	slices.SortFunc(roomSizeList, func(a, b *Room) int {
		return a.size - b.size
	})

	var p *Room
	for i, roomPtr := range roomIdList {
		room := roomPtr
		room.idx = i
		if p != nil {
			p.next = room
			room.prev = p
		}
		p = room
	}

	queryList := make([]*Queries, 0, len(queries))
	for idx, query := range queries {
		queryList = append(queryList, &Queries{
			idx:  idx,
			id:   query[0],
			size: query[1],
		})
	}
	slices.SortFunc(queryList, func(a, b *Queries) int {
		return a.size - b.size
	})

	sizeIdx := 0
	ret := make([]int, len(queries))
	for _, query := range queryList {
		preferred := query.id
		minSize := query.size
		for ; sizeIdx < n; sizeIdx++ {
			minSizeRoom := roomSizeList[sizeIdx]
			if minSizeRoom.size >= minSize {
				break
			}

			idx := minSizeRoom.idx
			prev := minSizeRoom.prev
			next := minSizeRoom.next
			if prev != nil {
				prev.next = next
				(*roomIdList[idx]).real = prev
			}
			if next != nil {
				next.prev = prev
				(*roomIdList[idx]).real = next
			}
		}

		if sizeIdx == n {
			ret[query.idx] = -1
			continue
		}

		// 拷贝slices.BinarySearchFunc改了一下
		i := binarySearchFunc(roomIdList, preferred, func(room *Room, id int) int {
			return room.RealId() - id
		})
		if i == n {
			i = i - 1
		}

		room := *roomIdList[i]
		if room.RealId() == preferred {
			ret[query.idx] = room.RealId()
		}

		if room.RealId() > preferred {
			p = room.prev
		} else {
			p = room.next
		}

		if p == nil {
			ret[query.idx] = room.RealId()
		} else {
			a := abs(room.RealId() - preferred)
			b := abs(p.RealId() - preferred)
			if a < b {
				ret[query.idx] = room.RealId()
			} else if b < a {
				ret[query.idx] = p.RealId()
			} else if room.RealId() < p.RealId() {
				ret[query.idx] = room.RealId()
			} else {
				ret[query.idx] = p.RealId()
			}
		}
	}
	return ret
}

func main() {
	//fmt.Println(closestRoom([][]int{{23, 22}, {6, 20}, {15, 6}, {22, 19}, {2, 10}, {21, 4}, {10, 18}, {16, 1}, {12, 7}, {5, 22}}, [][]int{{21, 6}}))
	//fmt.Println(closestRoom([][]int{{2, 2}, {1, 2}, {3, 2}}, [][]int{{3, 1}, {3, 3}, {5, 2}}))
	//fmt.Println(closestRoom([][]int{{1, 4}, {2, 3}, {3, 5}, {4, 1}, {5, 2}}, [][]int{{2, 3}, {2, 4}, {2, 5}}))
}
