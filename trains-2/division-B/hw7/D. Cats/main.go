package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type event struct {
	id int
	x  int
	t  int
}

const (
	begin = iota
	cat
	end
)

type line struct {
	l int
	r int
}

func main() {
	rb := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(rb, &n, &m)

	events := make([]event, 0, n+m*2)
	for range n {
		var x int
		fmt.Fscan(rb, &x)
		events = append(events, event{
			x: x,
			t: cat,
		})
	}
	lines := make([]line, 0, m)
	for i := range m {
		var l, r int
		fmt.Fscan(rb, &l, &r)
		events = append(events,
			event{id: i, x: l, t: begin},
			event{id: i, x: r, t: end})
		lines = append(lines, line{l: l, r: r})

	}
	sort.Slice(events, func(i, j int) bool {
		if events[i].x == events[j].x {
			return events[i].t < events[j].t
		}
		return events[i].x < events[j].x
	})

	catsAtBegin := make(map[int]int, m)
	catsAtEnd := make(map[int]int, m)
	nowCats := 0
	for _, e := range events {
		switch e.t {
		case cat:
			nowCats++
		case begin:
			catsAtBegin[e.x] = nowCats
		case end:
			catsAtEnd[e.x] = nowCats
		}
	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, l := range lines {
		fmt.Fprintln(w, catsAtEnd[l.r]-catsAtBegin[l.l])
	}
}
