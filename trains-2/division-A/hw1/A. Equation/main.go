package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	switch {
	case a == 0 && b == 0:
		fmt.Println("INF")
	case b == 0 && d != 0:
		fmt.Println(0)
	case a != 0 && c != 0 && -b/a == -d/c:
		fmt.Println("NO")
	case a != 0 && b != 0:
		if b%a == 0 {
			fmt.Println(-b / a)
		} else {
			fmt.Println("NO")
		}
	}
}
