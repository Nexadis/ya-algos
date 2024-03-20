package main

import (
	"fmt"
)

func main() {
	var size, elem, x int
	fmt.Scan(&size)
	list := make([]int, 0, size)
	for i := 0; i < size; i++ {
		fmt.Scan(&elem)
		list = append(list, elem)
	}
	fmt.Scan(&x)
	nearest := solution(list, x)
	fmt.Println(nearest)
}

func solution(list []int, x int) int {
	num := list[0]
	diff := difference(list[0], x)
	for _, v := range list {
		d := difference(v, x)
		if d < diff {
			diff = d
			num = v
		}
	}
	return num
}

func difference(a, b int) int {
	return max(a, b) - min(a, b)
}
