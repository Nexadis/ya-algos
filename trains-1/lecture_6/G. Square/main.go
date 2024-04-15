package main

import "fmt"

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() int64 {
	var n, m, t int64

	fmt.Scan(&n, &m, &t)

	l := int64(0)
	r := max(n, m)

	check := func(width int64) bool {
		if width > (min(n, m)-1)/2 {
			return false
		}
		spend := (m - width + n - width) * width * 2
		return spend <= t
	}

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
