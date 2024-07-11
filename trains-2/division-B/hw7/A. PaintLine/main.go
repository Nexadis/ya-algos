package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type action struct {
	x     int
	atype int
}

const (
	end   = -1
	start = 1
)

func main() {
	var n int

	rb := bufio.NewReader(os.Stdin)

	fmt.Fscan(rb, &n)
	actions := make([]action, 0, n*2)
	for range n {
		var s, e int
		fmt.Fscan(rb, &s, &e)
		actions = append(actions,
			action{x: s, atype: start},
			action{x: e, atype: end},
		)
	}
	sort.Slice(actions, func(i, j int) bool {
		if actions[i].x == actions[j].x {
			return actions[i].atype < actions[j].atype
		}
		return actions[i].x < actions[j].x
	})

	l := 0
	cnt := 0
	begin := actions[0].x
	for _, a := range actions {
		if a.atype == start {
			if cnt == 0 {
				begin = a.x
			}
			cnt++
		}
		if a.atype == end {
			cnt--
			if cnt == 0 {
				l += a.x - begin
			}
		}
	}
	fmt.Println(l)
}
