package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	end = iota
	begin
)

type block struct {
	id    int
	area  int
	z     int
	event int
}

type ByZ []block

func (a ByZ) Len() int      { return len(a) }
func (a ByZ) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByZ) Less(i, j int) bool {
	if a[i].z == a[j].z {
		return a[i].event < a[j].event
	}
	return a[i].z < a[j].z
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, w, l int
	fmt.Fscan(r, &n, &w, &l)
	needArea := w * l
	blocks := make([]block, 0, n*2)
	for i := 1; i < n+1; i++ {
		var x1, y1, z1, x2, y2, z2 int
		fmt.Fscan(r, &x1, &y1, &z1)
		fmt.Fscan(r, &x2, &y2, &z2)
		area := (x2 - x1) * (y2 - y1)

		blocks = append(blocks,
			block{id: i, area: area, z: z1, event: begin},
			block{id: i, area: area, z: z2, event: end},
		)
	}

	area := 0
	bCnt := 0
	minCnt := n + 1
	sort.Sort(ByZ(blocks))
	for _, b := range blocks {
		if b.event == begin {
			bCnt++
			area += b.area
		} else {
			bCnt--
			area -= b.area
		}
		if area == needArea && bCnt < minCnt {
			minCnt = bCnt
		}
	}
	wb := bufio.NewWriter(os.Stdout)
	defer wb.Flush()

	if minCnt == n+1 {
		fmt.Fprintln(wb, "NO")
		return
	}
	bestSeq := make([]int, 0, minCnt)
	nowBlocks := make(map[int]struct{}, minCnt)

	for _, b := range blocks {
		if b.event == begin {
			nowBlocks[b.id] = struct{}{}
			bCnt++
			area += b.area
		} else {
			delete(nowBlocks, b.id)
			bCnt--
			area -= b.area
		}
		if area == needArea && bCnt == minCnt {
			for id := range nowBlocks {
				bestSeq = append(bestSeq, id)
			}
			break
		}
	}
	fmt.Fprintln(wb, "YES")
	fmt.Fprintln(wb, minCnt)
	for _, id := range bestSeq {
		fmt.Fprintln(wb, id)
	}
}
