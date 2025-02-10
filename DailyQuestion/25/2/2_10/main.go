package main

import (
	"fmt"
)

func main() {
	//ret := catMouseGame([][]int{
	//	{2, 5},
	//	{3},
	//	{0, 4, 5},
	//	{1, 4, 5},
	//	{2, 3},
	//	{0, 2, 3},
	//})

	//ret := catMouseGame([][]int{
	//	{5, 6},
	//	{3, 4},
	//	{6},
	//	{1, 4, 5},
	//	{1, 3, 5},
	//	{0, 3, 4, 6},
	//	{0, 2, 5},
	//})

	ret := catMouseGame([][]int{
		{6},
		{4},
		{9},
		{5},
		{1, 5},
		{3, 4, 6},
		{0, 5, 10},
		{8, 9, 10},
		{7},
		{2, 7},
		{6, 7},
	})
	fmt.Println(ret)
}

func toStr(typ ResultType) string {
	switch typ {
	case Pending:
		return "待定"
	case CatWin:
		return "猫赢"
	case MouseWin:
		return "鼠赢"
	}
	return ""
}

func catMouseGame(graph [][]int) int {
	game := NewGame(graph)
	result := game.mouseRun(Status{mouse: 1, cat: 2})

	//var strings []string
	//for status, ret := range game.mouse {
	//	strings = append(strings, fmt.Sprintf("鼠行动 Mouse %d Cat %d  Type %s\n", status.mouse, status.cat, toStr(ret.Type)))
	//}
	//for status, ret := range game.cat {
	//	strings = append(strings, fmt.Sprintf("猫行动 Cat %d Mouse %d Type %s\n", status.cat, status.mouse, toStr(ret.Type)))
	//}
	//slices.Sort(strings)
	//for _, s := range strings {
	//	fmt.Print(s)
	//}

	isContinue := true
	for isContinue {
		isContinue = false
		for _, ret := range game.mouse {
			if ret.Type == Pending {
				ret.tryRemovePending()
				if ret.Type != Pending {
					isContinue = true
				}
			}
		}

		for _, ret := range game.cat {
			if ret.Type == Pending {
				ret.tryRemovePending()
				if ret.Type != Pending {
					isContinue = true
				}
			}
		}
	}

	//fmt.Println("")
	//strings = nil
	//for status, ret := range game.mouse {
	//	strings = append(strings, fmt.Sprintf("鼠行动 Mouse %d Cat %d  Type %s\n", status.mouse, status.cat, toStr(ret.Type)))
	//}
	//for status, ret := range game.cat {
	//	strings = append(strings, fmt.Sprintf("猫行动 Cat %d Mouse %d Type %s\n", status.cat, status.mouse, toStr(ret.Type)))
	//}
	//slices.Sort(strings)
	//for _, s := range strings {
	//	fmt.Print(s)
	//}

	if result.IsPending() {
		return int(Peace)
	}
	return int(result.Type)
}

type ResultType int

type Result struct {
	Animal      int // 1 猫 2 鼠
	Type        ResultType
	PendingList []*Result // 待定的列表，在该结果求出具体值时，将列表中的结果设值
	References  []Status
}

const (
	AnimalCat   = 1
	AnimalMouse = 2
)

func (s *Result) IsPending() bool {
	if s.Type != Pending {
		return false
	}

	if len(s.PendingList) > 0 {
		s.tryRemovePending()
	}

	return s.Type == Pending
}

func (s *Result) tryRemovePending() {
	var ret ResultType
	if s.Animal == AnimalCat {
		ret = MouseWin
	} else {
		ret = CatWin
	}

	for _, result := range s.PendingList {
		if result.Type == Pending {
			ret = Pending
			continue
		}

		if s.Animal == AnimalCat && result.Type == CatWin {
			s.Type = CatWin
			return
		}

		if s.Animal == AnimalMouse && result.Type == AnimalMouse {
			s.Type = AnimalMouse
			return
		}
	}

	s.Type = ret
}

func NewResult(animal int) *Result {
	return &Result{Animal: animal, Type: Pending}
}

const (
	Peace    ResultType = 0
	MouseWin ResultType = 1
	CatWin   ResultType = 2
	Pending  ResultType = 3
)

type Status struct {
	cat   int
	mouse int
}

type Game struct {
	graph [][]int
	mouse map[Status]*Result // mouse的回合，这个状态下得到的结果
	cat   map[Status]*Result // cat的回合，这个状态下得到的结果
}

func NewGame(graph [][]int) *Game {
	return &Game{
		graph: graph,
		mouse: make(map[Status]*Result),
		cat:   make(map[Status]*Result),
	}
}

func (s *Game) setMouseStatus(status Status, typ ResultType) *Result {
	result := s.mouse[status]
	result.Type = typ
	return result
}

func (s *Game) setCatStatus(status Status, typ ResultType) *Result {
	result := s.cat[status]
	result.Type = typ
	return result
}

func (s *Game) mouseRun(status Status) *Result {
	ret := CatWin                            // 先设为最差情况
	s.mouse[status] = NewResult(AnimalMouse) // 标记为待定
	var pendingList []*Result

	path := s.graph[status.mouse]
	for _, dstPoint := range path {
		if dstPoint == 0 {
			return s.setMouseStatus(status, MouseWin)
		}

		if dstPoint == status.cat {
			continue // 走到了猫的位置，必输
		}

		dstStatus := Status{mouse: dstPoint, cat: status.cat}
		result, ok := s.cat[dstStatus]
		if !ok {
			result = s.catRun(dstStatus)
		}

		if result.IsPending() {
			ret = Pending // 存在平局的可能，将结果设置为待定
			pendingList = append(pendingList, result)
			continue // 继续寻找更好的结果
		}

		if result.Type == MouseWin {
			return s.setMouseStatus(status, MouseWin)
		}
	}

	if ret == Pending { // 如果最终结果为待定
		s.mouse[status].PendingList = pendingList
		return s.mouse[status]
	}

	return s.setMouseStatus(status, ret)
}

func (s *Game) catRun(status Status) *Result {
	ret := MouseWin                      // 先设为最差情况
	s.cat[status] = NewResult(AnimalCat) // 标记为待定
	var pendingList []*Result

	path := s.graph[status.cat]
	for _, dstPoint := range path {
		if dstPoint == 0 {
			continue // 不可以走到0
		}

		if dstPoint == status.mouse {
			return s.setCatStatus(status, CatWin)
		}

		dstStatus := Status{mouse: status.mouse, cat: dstPoint}
		result, ok := s.mouse[dstStatus]
		if !ok {
			result = s.mouseRun(dstStatus)
		}

		if result.IsPending() {
			ret = Pending // 存在平局的可能，将结果设置为待定
			pendingList = append(pendingList, result)
			continue // 继续寻找更好的结果
		}

		if result.Type == CatWin {
			return s.setCatStatus(status, CatWin)
		}
	}

	if ret == Pending { // 如果最终结果为待定
		s.cat[status].PendingList = pendingList
		return s.cat[status]
	}

	return s.setCatStatus(status, ret)
}
