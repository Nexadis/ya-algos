package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	ans := solution(a, b, c, d, e)
	fmt.Println(ans)
}

func solution(a, b, c, d, e int) string {
	a, b = sort(a, b)
	b, _ = sort(b, c)
	a, b = sort(a, b)
	d, e = sort(d, e)
	if a <= d && b <= e {
		return "YES"
	}
	return "NO"
}

func sort(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}
