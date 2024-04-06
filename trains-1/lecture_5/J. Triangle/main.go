package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	triangles := solution()
	fmt.Println(triangles)
}

type point struct {
	x int
	y int
}

func solution() int {
	points := input(os.Stdin)

	sum := 0
	for i, p1 := range points {
		usedvectors := map[point]struct{}{}
		neighbours := make([]int, 0, len(points)-i)

		for _, p2 := range points {
			vec := point{}
			vec.x = p1.x - p2.x
			vec.y = p1.y - p2.y
			veclen := vec.x*vec.x + vec.y*vec.y
			neighbours = append(neighbours, veclen)
			if _, ok := usedvectors[vec]; ok {
				sum--
			}
			vec.x = -vec.x
			vec.y = -vec.y
			usedvectors[vec] = struct{}{}
		}

		slices.Sort(neighbours)
		r := 0
		for l := range neighbours {
			for r < len(neighbours) && neighbours[l] == neighbours[r] {
				r++
			}
			sum += r - l - 1
		}

	}
	return sum
}

func input(r io.Reader) []point {
	rd := bufio.NewReader(r)
	var n int
	fmt.Fscan(rd, &n)
	points := make([]point, 0, n)
	for i := 0; i < n; i++ {
		var p point
		fmt.Fscan(rd, &p.x, &p.y)
		points = append(points, p)
	}
	return points
}
