package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Scan(&a, &b, &c, &d, &e)
	ans := solution(a, b, c, d, e)
	fmt.Println(ans)
}

func solution(a, b, c, d, e int) string {
	if a <= d {
		if b <= e || c <= e {
			return "YES"
		}
	}
	if b <= d {
		if a <= e || c <= e {
			return "YES"
		}
	}
	if c <= d {
		if a <= e || b <= e {
			return "YES"
		}
	}
	if a <= e {
		if b <= d || c <= d {
			return "YES"
		}
	}
	if b <= e {
		if a <= d || c <= d {
			return "YES"
		}
	}
	if c <= e {
		if a <= d || b <= d {
			return "YES"
		}
	}
	return "NO"
}
