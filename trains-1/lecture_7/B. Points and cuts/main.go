package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	points, lenPoints := input(os.Stdin)
	includes := solution(points, lenPoints)
	out(os.Stdout, includes)
}

func out(w io.Writer, includes []int) {
	wb := bufio.NewWriter(w)
	defer wb.Flush()
	for _, in := range includes {
		fmt.Fprint(wb, in, " ")
	}
}

func input(r io.Reader) ([]Point, int) {
	rb := bufio.NewReader(r)
	var n, m int
	fmt.Fscan(rb, &n, &m)
	points := make([]Point, 0, n*2+m)
	for i := 0; i < n; i++ {
		var x1, x2 int
		fmt.Fscan(rb, &x1, &x2)
		s := Point{what: start, x: min(x1, x2)}
		e := Point{what: end, x: max(x1, x2)}
		points = append(points, s, e)
	}
	for i := 0; i < m; i++ {
		p := Point{what: point, num: i}
		fmt.Fscan(rb, &p.x)
		points = append(points, p)
	}

	return points, m
}

func solution(points []Point, lenPoints int) []int {
	sort.Sort(ByPlace(points))
	includes := make([]int, lenPoints)
	nowInclude := 0
	for _, p := range points {
		switch p.what {
		case start:
			nowInclude += 1
		case end:
			nowInclude -= 1
		case point:
			includes[p.num] = nowInclude
		}
	}
	return includes
}

type kind int

const (
	start kind = iota
	point
	end
)

type Point struct {
	x    int
	what kind
	num  int
}

type ByPlace []Point

func (a ByPlace) Len() int      { return len(a) }
func (a ByPlace) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPlace) Less(i, j int) bool {
	if a[i].x == a[j].x {
		return a[i].what < a[j].what
	}
	return a[i].x < a[j].x
}
