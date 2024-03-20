package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	list := input(n)
	add, appendix := solution(0, list)
	if add == 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(add)
	for _, v := range appendix {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func input(n int) []int {
	var x int
	list := make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		list = append(list, x)
	}
	return list
}

func solution(start int, list []int) (int, []int) {
	left, right := start, len(list)-1
	for left <= right {
		if list[left] != list[right] {
			return solution(left+1, list)
		}
		left += 1
		right -= 1

	}
	return start, inverse(list[:start])
}

func inverse(list []int) []int {
	out := make([]int, len(list))
	for i, v := range list {
		out[len(list)-1-i] = v
	}
	return out
}
