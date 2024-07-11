package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type sector struct {
	id int
	r  float64
	t  int
	s  float64
	e  float64
}

const (
	finish = iota
	begin
)

type line struct {
	x float64
	t int
}

func main() {
	sectors := input()
	curSectors := make(map[int]sector, len(sectors))

	area := float64(0)
	prev := sectors[0].r
	for _, s := range sectors {
		if len(sectors)/2 == len(curSectors) {
			l := sumLen(curSectors, len(sectors)/2)
			area += l / 2 * (s.r*s.r - prev*prev)
		}
		prev = s.r

		if s.t == begin {
			curSectors[s.id] = s
		} else {
			delete(curSectors, s.id)
		}
	}
	fmt.Println(area)
}

func sumLen(sectors map[int]sector, sCnt int) (length float64) {
	if len(sectors) == 0 {
		return 0
	}
	sorted := make([]line, 0, 2*len(sectors))
	for _, s := range sectors {
		if s.s > s.e {
			sorted = append(sorted,
				line{x: 0, t: begin},
				line{x: 2 * math.Pi, t: finish},
			)
		}
		sorted = append(sorted,
			line{x: s.s, t: begin},
			line{x: s.e, t: finish},
		)
	}
	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].x == sorted[j].x {
			return sorted[i].t > sorted[j].t
		}
		return sorted[i].x < sorted[j].x
	})

	l := float64(0)
	cnt := 0
	start := float64(0)
	for _, s := range sorted {
		if s.t == begin {
			cnt++
			if cnt == sCnt {
				start = s.x
			}
			continue
		} else {
			if cnt == sCnt {
				l += s.x - start
			}
			cnt--
		}
	}
	return l
}

func input() []sector {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	sectors := make([]sector, 0, n*2)
	for i := range n {
		var r1, r2, f1, f2 float64
		fmt.Fscan(r, &r1, &r2, &f1, &f2)
		sectors = append(sectors,
			sector{id: i, r: r1, t: begin, s: f1, e: f2},
			sector{id: i, r: r2, t: finish, s: f1, e: f2},
		)
	}

	sort.Slice(sectors, func(i, j int) bool {
		if sectors[i].r == sectors[j].r {
			return sectors[i].t < sectors[j].t
		}
		return sectors[i].r < sectors[j].r
	})
	return sectors
}
