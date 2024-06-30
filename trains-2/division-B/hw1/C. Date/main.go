package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	x, y = max(x, y), min(x, y)
	if x > 12 && y <= 12 || x == y {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
