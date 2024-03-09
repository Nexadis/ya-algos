package main

import (
	"fmt"
)

func main() {
	var h1, w1, h2, w2 int
	fmt.Scan(&h1, &w1, &h2, &w2)

	h, w := solution(h1, w1, h2, w2)
	fmt.Println(h, w)
}

func solution(h1, w1, h2, w2 int) (int, int) {
	m1 := (h1 + h2) * max(w1, w2)
	m2 := (w1 + w2) * max(h1, h2)
	m3 := (h1 + w2) * max(w1, h2)
	m4 := (w1 + h2) * max(h1, w2)
	m := min(m1, m2, m3, m4)
	switch m {
	case m1:
		return h1 + h2, max(w1, w2)
	case m2:
		return w1 + w2, max(h1, h2)
	case m3:
		return h1 + w2, max(w1, h2)
	case m4:
		return w1 + h2, max(h1, w2)
	}
	return 0, 0
}
