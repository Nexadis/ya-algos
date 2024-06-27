package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	close = iota
	open
)

type event struct {
	id  int
	t   int
	min int
}

type ByTime []event

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool {
	if a[i].min == a[j].min {
		return a[i].t < a[j].t
	}
	return a[i].min < a[j].min
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	events := make([]event, 0, n)
	for i := 0; i < n; i++ {
		var oHour, oMin, cHour, cMin int
		fmt.Fscan(r, &oHour, &oMin, &cHour, &cMin)

		if (oHour-cHour)*60+(oMin-cMin) >= 0 {
			events = append(events,
				event{id: i, t: open, min: 0},
				event{id: i, t: close, min: 24 * 60},
			)
		}
		events = append(events,
			event{id: i, t: open, min: oHour*60 + oMin},
			event{id: i, t: close, min: cHour*60 + cMin})
	}

	sort.Sort(ByTime(events))

	sumTime := 0
	opened := 0
	last := 0
	for _, e := range events {
		if e.t == open {
			opened++
			last = e.min
		} else {
			if opened == n {
				sumTime += e.min - last
			}
			opened--
		}
	}

	fmt.Println(sumTime)
}
