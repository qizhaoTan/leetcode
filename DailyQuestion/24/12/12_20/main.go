package main

import (
	"slices"
)

func maxCommonDivisors(numbers ...int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}

	x := numbers[0]
	for _, y := range numbers[1:] {
		x = maxCommonDivisor(x, y)
	}

	return x
}

func maxCommonDivisor(x, y int) int {
	if x == y {
		return x
	}

	if x < y {
		return maxCommonDivisors(y, x)
	}

	if y == 0 {
		return x
	}

	m := x % y
	return maxCommonDivisors(y, m)
}

func calcCharNumbers(s string) []int {
	arr := make([]int, 26)
	for _, c := range s {
		arr[c-'a']++
	}
	return arr
}

func isMeet(s string, n int) bool {
	src := calcCharNumbers(s[:n])
	for i := n; i < len(s); i += n {
		if !slices.Equal(calcCharNumbers(s[i:i+n]), src) {
			return false
		}
	}
	return true
}

func minAnagramLength(s string) int {
	arr := calcCharNumbers(s)

	var numbers []int
	for _, num := range arr {
		numbers = append(numbers, num)
	}

	n := maxCommonDivisors(numbers...)
	for i := n; i >= 2; i-- {
		if n%i != 0 {
			continue
		}

		m := len(s) / i
		if isMeet(s, m) {
			return m
		}
	}

	return len(s)
}

func main() {
	println(minAnagramLength("jjj"))
	println(minAnagramLength("abba"))
	println(minAnagramLength("cdef"))
	println(minAnagramLength("abbabbaa"))
	println(minAnagramLength("pqqppqpqpq"))
}
