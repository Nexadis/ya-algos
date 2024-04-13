package main

import "fmt"

func main() {
	secs := solution()
	fmt.Println(secs)
}

func solution() int {
	var n, x, y int
	fmt.Scan(&n, &x, &y)
	l := 0
	r := n * max(x, y)

	slow, fast := max(x, y), min(x, y)

	check := func(seconds int) bool {
		// с одного листа нельзя сделать сразу 2 копии
		// поэтому делаем первую копию на самом быстром принтере
		if seconds < fast {
			return false
		}
		seconds -= fast
		// одна копия в начале на быстром принтере
		return seconds/slow+seconds/fast >= n-1
	}
	for l < r {
		m := (l + r) / 2
		if check(m) {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
