package __26

/*
你有一个只支持单个标签页的 浏览器 ，最开始你浏览的网页是 homepage ，你可以访问其他的网站 url ，也可以在浏览历史中后退 steps 步或前进 steps 步。

请你实现 BrowserHistory 类：

BrowserHistory(string homepage) ，用 homepage 初始化浏览器类。
void visit(string url) 从当前页跳转访问 url 对应的页面  。执行此操作会把浏览历史前进的记录全部删除。
string back(int steps) 在浏览历史中后退 steps 步。如果你只能在浏览历史中后退至多 x 步且 steps > x ，那么你只后退 x 步。请返回后退 至多 steps 步以后的 url 。
string forward(int steps) 在浏览历史中前进 steps 步。如果你只能在浏览历史中前进至多 x 步且 steps > x ，那么你只前进 x 步。请返回前进 至多 steps步以后的 url 。
*/

type BrowserHistory struct {
	arr []string
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		arr: []string{homepage},
	}
}

func (s *BrowserHistory) Visit(url string) {
	s.arr = append(s.arr, url)
	s.arr = s.arr[:len(s.arr):len(s.arr)]
}

func (s *BrowserHistory) Back(steps int) string {
	if steps >= len(s.arr) {
		steps = len(s.arr) - 1
	}

	s.arr = s.arr[:len(s.arr)-steps]
	return s.arr[len(s.arr)-1]
}

func (s *BrowserHistory) Forward(steps int) string {
	if steps > cap(s.arr)-len(s.arr) {
		steps = cap(s.arr) - len(s.arr)
	}

	s.arr = s.arr[:len(s.arr)+steps]
	return s.arr[len(s.arr)-1]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
