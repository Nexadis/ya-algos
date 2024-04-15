package main

import "fmt"

func main() {
	fives := solution()
	fmt.Println(fives)
}

func solution() int {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	l := 0
	r := a + b + c
	check := func(fives int) bool {
		return 10*(a*2+b*3+c*4+fives*5) >= 35*(a+b+c+fives)
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
