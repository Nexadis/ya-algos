package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if c < 0 {
		fmt.Println("NO SOLUTION")
		return
	}
	if a == 0 {
		if b == c*c {
			fmt.Println("MANY SOLUTIONS")
			return
		}
		fmt.Println("NO SOLUTION")
		return
	}
	if (c*c-b)%a != 0 {
		fmt.Println("NO SOLUTION")
		return
	}
	x := (c*c - b) / a
	fmt.Println(x)
}
