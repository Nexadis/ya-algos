package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type line struct {
	l int
	r int
}

func main() {
	var m int
	rb := bufio.NewReader(os.Stdin)
	fmt.Fscan(rb, &m)

	lines := make([]line, 0, 1000*2)

	for {
		var l, r int
		fmt.Fscan(rb, &l, &r)
		if l == 0 && r == 0 {
			break
		}
		if l > m || r < 0 {
			continue
		}
		ln := line{l: l, r: r}
		lines = append(lines, ln)
	}

	if len(lines) == 0 {
		fmt.Println("No solution")
		return
	}

	sort.Slice(lines, func(i, j int) bool {
		if lines[i].l == lines[j].l {
			return lines[i].r < lines[j].r
		}
		return lines[i].l < lines[j].l
	})

	need := make([]line, 0, len(lines))
	nowEnd := 0
	bestEnd := 0
	best := 0

	for i := 0; i < len(lines); i++ {
		l := lines[i]
		if l.l > nowEnd {
			need = append(need, lines[best])
			nowEnd = bestEnd
			if nowEnd >= m {
				break
			}
		}
		if l.l <= nowEnd && l.r > bestEnd {
			bestEnd = l.r
			best = i
		}
	}

	if nowEnd < m {
		nowEnd = bestEnd
		need = append(need, lines[best])
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	if nowEnd < m {
		fmt.Fprintln(w, "No solution")
		return
	}

	sort.Slice(need, func(i, j int) bool {
		return need[i].l < need[j].l
	})
	fmt.Fprintln(w, len(need))

	for _, l := range need {
		fmt.Fprintln(w, l.l, l.r)
	}
}
