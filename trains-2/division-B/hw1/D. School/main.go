package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	houses := make([]int, n)
	minDistance := 0
	fmt.Fscan(r, &houses[0])
	for i := 1; i < n; i++ {
		fmt.Fscan(r, &houses[i])
		minDistance += houses[i] - houses[0]
	}

	distance := minDistance
	coord := 0
	for i := 1; i < n; i++ {
		diff := houses[i] - houses[i-1]
		distance = distance + i*diff - (n-i)*diff
		if distance <= minDistance {
			minDistance = distance
			coord = i
		}
	}
	fmt.Println(houses[coord])
}
