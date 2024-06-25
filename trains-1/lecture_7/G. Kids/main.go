package main

import (
	"bufio"
	"fmt"
	"os"
)

type helper struct {
	id     int
	t      int
	z      int
	y      int
	filled int
}

func main() {
	var m, n int

	r := bufio.NewReader(os.Stdin)

	fmt.Fscan(r, &m, &n)

	helpers := make([]helper, n)
	for i := 0; i < n; i++ {
		var t, z, y int
		fmt.Fscan(r, &t, &z, &y)
		helpers[i] = helper{id: i, t: t, z: z, y: y}
	}

	for i := 0; i < m; i++ {
		nowFill := 0
		minTime := whenFill(helpers[0])
		for j := range helpers {
			fillTime := whenFill(helpers[j])
			if fillTime < minTime {
				minTime = fillTime
				nowFill = j
			}
		}
		h := helpers[nowFill]
		h.filled++
		helpers[nowFill] = h
	}

	maxTime := 0
	for _, h := range helpers {
		fillTime := spentTime(h)
		if fillTime > maxTime && h.filled > 0 {
			maxTime = fillTime
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fprintln(w, maxTime)
	for _, h := range helpers {
		fmt.Fprint(w, h.filled, " ")
	}
}

func spentTime(h helper) int {
	if h.filled == 0 {
		return 0
	}
	fillTime := h.t * h.filled
	restTimes := h.filled / h.z
	if h.filled%h.z == 0 && restTimes > 0 {
		restTimes -= 1
	}
	return fillTime + restTimes*h.y
}

func whenFill(h helper) int {
	if h.filled == 0 {
		return h.t
	}

	fillTime := h.filled*h.t + h.t
	restTimes := h.filled / h.z
	return fillTime + restTimes*h.y
}
