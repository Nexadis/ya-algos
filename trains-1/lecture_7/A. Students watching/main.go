package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	tables, students := input(os.Stdin)
	watching := solution(tables)
	fmt.Println(students - watching)
}

func solution(watches []watch) int {
	watching := 0
	sorted := SortBy(watches)
	sort.Sort(sorted)
	watchers := 0

	for i, w := range sorted {
		if watchers > 0 {
			prev := sorted[i-1]
			watching += w.table - prev.table
		}
		if w.event == Start {
			watchers += 1
		}
		if w.event == End {
			watchers -= 1
			if watchers == 0 {
				watching++
			}
		}

	}
	return watching
}

func input(r io.Reader) ([]watch, int) {
	rb := bufio.NewReader(r)
	var n, m int
	fmt.Fscan(rb, &n, &m)
	watches := make([]watch, m*2)
	for i := 0; i < len(watches); i += 2 {
		fmt.Fscan(rb, &watches[i].table, &watches[i+1].table)
		watches[i].event = Start
		watches[i+1].event = End
	}
	return watches, n
}

type event int

const (
	Start event = iota
	End
)

type watch struct {
	table int
	event event
}

type SortBy []watch

func (a SortBy) Len() int      { return len(a) }
func (a SortBy) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {
	if a[i].table == a[j].table {
		return a[i].event < a[j].event
	}
	return a[i].table < a[j].table
}
