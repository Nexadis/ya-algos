package main

import "fmt"

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() int {
	var n int
	fmt.Scan(&n)
	shoots := make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		shoots[x] = struct{}{}
	}
	return len(shoots)
}
