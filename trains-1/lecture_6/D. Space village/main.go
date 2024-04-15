package main

import "fmt"

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() int64 {
	var n, a, b, w, h int64
	fmt.Scan(&n, &a, &b, &w, &h)
	check1 := func(m int64) bool {
		return (w/(a+2*m))*(h/(b+2*m)) >= n
	}
	check2 := func(m int64) bool {
		return (w/(b+2*m))*(h/(a+2*m)) >= n
	}
	ans1 := rightSearch(0, max(w, h), check1)
	ans2 := rightSearch(0, max(w, h), check2)
	return max(ans1, ans2)
}

func rightSearch(from, to int64, check func(m int64) bool) int64 {
	l := from
	r := to

	for l < r {
		m := (l + r + 1) / 2
		if check(m) {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}
