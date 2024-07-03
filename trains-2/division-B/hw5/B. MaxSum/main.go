package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	rb := bufio.NewReader(os.Stdin)
	fmt.Fscan(rb, &n)
	sums := make([]int, n)
	var a int
	fmt.Fscan(rb, &a)
	sum := a
	sums[0] = sum
	for i := 1; i < n; i++ {
		var a int
		fmt.Fscan(rb, &a)
		if sum < 0 && a > 0 {
			sum = a
		} else {
			sum += a
		}

		sums[i] = sum
	}
	maxSum := sums[0]
	for _, s := range sums {
		maxSum = max(maxSum, s)
	}
	fmt.Println(maxSum)
}
