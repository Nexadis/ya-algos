package main

import "fmt"

func main() {
	var a, b, n, m int
	fmt.Scan(&a, &b, &n, &m)

	min, max := solution(a, b, n, m)

	if min == -1 {
		fmt.Println(-1)
		return
	}
	fmt.Println(min, max)
}

func getMin(wait, trains int) int {
	return wait*(trains-1) + trains
}

func getMax(wait, trains int) int {
	return wait*(trains+1) + trains
}

func solution(a, b, n, m int) (int, int) {
	min1 := getMin(a, n)
	max1 := getMax(a, n)
	min2 := getMin(b, m)
	max2 := getMax(b, m)
	if min2 > max1 || min1 > max2 {
		return -1, -1
	}
	return max(min1, min2), min(max1, max2)
}
