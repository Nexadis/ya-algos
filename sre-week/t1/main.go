package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type session struct {
	t       int
	isStart bool
}

type ByStart []session

func (s ByStart) Len() int      { return len(s) }
func (s ByStart) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByStart) Less(i, j int) bool {
	if s[i].t == s[j].t {
		return s[i].isStart
	}
	return s[i].t < s[j].t
}

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	sessions := make([]session, n*2)

	for i := 0; i < n*2; i += 2 {
		s := session{isStart: true}
		fmt.Fscan(r, &s.t)
		sessions[i] = s

		s.isStart = false
		fmt.Fscan(r, &s.t)
		sessions[i+1] = s
	}
	sort.Sort(ByStart(sessions))
	maxActive := 0
	active := 0
	moment := 0
	for _, s := range sessions {
		if s.isStart {
			active++
		}

		if active > maxActive {
			maxActive = active
			moment = s.t
		}

		if !s.isStart {
			active--
		}
	}
	fmt.Println(moment)
}
