package main

import "fmt"

type TrieTree struct {
	arr [26]*TrieTree
}

func NewTrieTree(words []string) *TrieTree {
	tire := &TrieTree{}
	for _, word := range words {
		tire.fill(word)
	}
	return tire
}

func (s *TrieTree) fill(word string) {
	trie := s
	for _, c := range word {
		idx := c - 'a'
		if trie.arr[idx] == nil {
			trie.arr[idx] = &TrieTree{}
		}

		trie = trie.arr[idx]
	}
}

func to(trie *TrieTree, target string) []int {
	ret := make([]int, len(target))
	for i := range target {
		ret[i] = calc(trie, target[i:]) + i
	}
	return ret
}

func calc(trie *TrieTree, target string) int {
	for i, c := range target {
		idx := c - 'a'
		if trie.arr[idx] == nil {
			return i
		}

		trie = trie.arr[idx]
	}
	return len(target)
}

func minValidStrings(words []string, target string) int {
	trie := NewTrieTree(words)
	s := to(trie, target)       // j = s[i] 表示在i这个位置，最长可以到达j这个位置
	dp := make([]int, len(s)+1) // 动态规划 dp[i] 表示从0到i至少需要多少次
	for idx, val := range s {
		if idx != 0 && dp[idx] == 0 { // 如果发现到达dp[idx]==0，说明无法到达上一个位置，那么可以判定为不可能满足条件
			return -1
		}

		step := dp[idx] + 1
		for i := idx + 1; i <= val; i++ {
			if dp[i] == 0 || dp[i] > step {
				dp[i] = step
			}
		}
	}

	ret := dp[len(s)]
	if ret == 0 { // 第一次提交漏此判断导致出错
		return -1
	}

	return ret
}

func main() {
	fmt.Println(minValidStrings([]string{"abc", "aaaaa", "bcdef"}, "aabcdabc"))
}
