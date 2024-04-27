package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	jumps := make([]int, n+1)
	jumps[1] = 1
	for i := 2; i < n+1; i++ {
		sum := 0
		start := i - k
		if i-k < 1 {
			start = 1
		}
		for j := start; j < i; j++ {
			sum += jumps[j]
		}
		jumps[i] = sum
	}
	fmt.Println(jumps[n])
}
