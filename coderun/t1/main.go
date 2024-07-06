package main

import (
	"fmt"
	"slices"
)

func main() {
	a := make([]int, 3)
	fmt.Scan(&a[0], &a[1], &a[2])
	slices.Sort(a)
	fmt.Println(a[1])
}
