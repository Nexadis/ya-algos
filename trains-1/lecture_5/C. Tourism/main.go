package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ans := solution()
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, a := range ans {
		fmt.Fprintln(w, a)
	}
}

func solution() []int {
	prefixUps, tracks := input()
	upsTrack := make([]int, len(tracks))
	for i, track := range tracks {
		var change int
		if track.end > track.start {
			change = prefixUps[track.end].up - prefixUps[track.start].up
		} else {
			change = prefixUps[track.start].down - prefixUps[track.end].down
		}
		upsTrack[i] = change

	}
	return upsTrack
}

func input() ([]diff, []track) {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n)
	prefixUps := make([]diff, n+1)
	prev := point{}
	fmt.Fscan(r, &prev.x, &prev.y)
	for i := 2; i <= n; i++ {
		next := point{}
		fmt.Fscan(r, &next.x, &next.y)
		change := next.y - prev.y
		if change > 0 {
			prefixUps[i].up = prefixUps[i-1].up + change
			prefixUps[i].down = prefixUps[i-1].down
		} else {
			prefixUps[i].up = prefixUps[i-1].up
			prefixUps[i].down = prefixUps[i-1].down - change
		}
		prev = next
	}

	fmt.Fscan(r, &m)
	tracks := make([]track, m)
	for i := range tracks {
		fmt.Fscan(r, &tracks[i].start, &tracks[i].end)
	}

	return prefixUps, tracks
}

type diff struct {
	up   int
	down int
}

type point struct {
	x int
	y int
}

type track struct {
	start int
	end   int
}
