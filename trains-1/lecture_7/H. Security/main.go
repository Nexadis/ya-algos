package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	var k int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fscan(r, &k)
	for i := 0; i < k; i++ {
		ans := test(r)
		fmt.Fprintln(w, ans)
	}
}

const (
	camein = iota
	cameout
)

type secman struct {
	id     int
	t      int
	action int
}

type ByTime []secman

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool {
	if a[i].t == a[j].t {
		return a[i].action < a[j].action
	}
	return a[i].t < a[j].t
}

func test(r io.Reader) string {
	var n int
	fmt.Fscan(r, &n)
	secs := make([]secman, 0, n*2)
	for i := 0; i < n; i++ {
		var in, out int
		fmt.Fscan(r, &in, &out)
		secs = append(secs,
			secman{id: i, t: in, action: camein},
			secman{id: i, t: out, action: cameout})
	}

	if isAccepted(secs, n) {
		return "Accepted"
	}
	return "Wrong Answer"
}

func isAccepted(secs []secman, n int) bool {
	sort.Sort(ByTime(secs))
	goodSecs := make(map[int]struct{}, n)
	nowWork := make(map[int]struct{}, n)

	prev := secman{t: -1}
	if secs[0].t != 0 {
		return false
	}

	for _, s := range secs {
		if len(nowWork) == 1 {
			for n := range nowWork {
				goodSecs[n] = struct{}{}
			}
		}
		if prev.t == s.t && prev.action == s.action {
			delete(goodSecs, prev.id)
			delete(goodSecs, s.id)
		}

		if s.action == camein {
			nowWork[s.id] = struct{}{}
		}
		if s.action == cameout {
			delete(nowWork, s.id)
		}
		if len(nowWork) == 0 && s.t != 10_000 {
			return false
		}
		prev = s

	}
	return len(goodSecs) == n
}

// ------
//     ------
//
// -----|----
//
// ------
// --------
//      -----
//
// ----
//  -----
//      -----
//     ------
//
//

func same(a, b secman) bool {
	return a.t == b.t && a.action == b.action
}
