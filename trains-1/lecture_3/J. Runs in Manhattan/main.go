package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data := input()
	ans := solution(data)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(out, len(ans))
	for _, p := range ans {
		fmt.Fprintln(out, p.x, p.y)
	}
	out.Flush()
}

type coordinate struct {
	x int
	y int
}

type rectangle struct {
	minp int
	minq int
	maxp int
	maxq int
}

type data struct {
	coords []coordinate
	t      int
	d      int
}

func input() data {
	var t, d, n int
	fmt.Scan(&t, &d, &n)
	coords := make([]coordinate, 0, n)
	for i := 0; i < n; i++ {
		c := coordinate{}
		fmt.Scan(&c.x, &c.y)
		coords = append(coords, c)
	}
	return data{coords, t, d}
}

func solution(d data) []coordinate {
	rect := rectangle{0, 0, 0, 0}
	for _, p := range d.coords {
		possibleRun := expand(rect, d.t)
		r := coord2Rect(p)
		possiblePos := expand(r, d.d)
		rect = intercept(possibleRun, possiblePos)
	}
	return rect2Coords(rect)
}

func intercept(r1, r2 rectangle) rectangle {
	r := rectangle{}
	r.maxq = min(r1.maxq, r2.maxq)
	r.minq = max(r1.minq, r2.minq)
	r.maxp = min(r1.maxp, r2.maxp)
	r.minp = max(r1.minp, r2.minp)
	return r
}

func expand(rect rectangle, delta int) rectangle {
	return rectangle{
		rect.minp - delta,
		rect.minq - delta,
		rect.maxp + delta,
		rect.maxq + delta,
	}
}

func coord2Rect(p coordinate) rectangle {
	return rectangle{
		p.x + p.y,
		p.x - p.y,
		p.x + p.y,
		p.x - p.y,
	}
}

func rect2Coords(r rectangle) []coordinate {
	coords := []coordinate{}
	for p := r.minp; p <= r.maxp; p++ {
		for q := r.minq; q <= r.maxq; q++ {
			if (p+q)%2 == 0 {
				x := (p + q) / 2
				y := p - x
				coords = append(coords, coordinate{x, y})

			}
		}
	}
	return coords
}
