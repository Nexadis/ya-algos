package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	uniqVars, variants := solution(os.Stdin)
	out(os.Stdout, uniqVars, variants)
}

const (
	start = iota
	end
)

type student struct {
	x  int
	id int
}

type ByX []student

func (a ByX) Len() int      { return len(a) }
func (a ByX) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByX) Less(i, j int) bool {
	return a[i].x < a[j].x
}

func solution(r io.Reader) (int, []int) {
	rb := bufio.NewReader(r)
	var n, d int
	fmt.Fscan(rb, &n, &d)

	students := make([]student, n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(rb, &x)
		students[i] = student{x: x, id: i}
	}

	sort.Sort(ByX(students))

	maxVars := 0

	variants := make([]int, n)
	prev := 0
	for i, s := range students {
		if students[i].x-students[prev].x > d {
			variants[s.id] = variants[students[prev].id]
			prev++
		} else {
			maxVars++
			variants[s.id] = maxVars
		}
	}

	return maxVars, variants
}

func out(w io.Writer, uniqVars int, variants []int) {
	wb := bufio.NewWriter(w)
	defer wb.Flush()
	fmt.Fprintln(wb, uniqVars)
	for _, v := range variants {
		fmt.Fprint(wb, v, " ")
	}
}
