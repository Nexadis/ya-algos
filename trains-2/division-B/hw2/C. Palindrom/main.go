package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	l := 0
	r := len(s) - 1
	cost := 0
	for l < r {
		if s[l] != s[r] {
			cost++
		}
		l++
		r--
	}
	fmt.Println(cost)
}
