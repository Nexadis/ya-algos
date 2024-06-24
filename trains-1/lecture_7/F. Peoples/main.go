package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"time"
)

const (
	death = iota
	birth
)

type man struct {
	t time.Time

	id    int
	event int
}

type ByTime []man

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool {
	if a[i].t.Equal(a[j].t) {
		return a[i].event < a[j].event
	}
	return a[i].t.Before(a[j].t)
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	people := readPeople(r, n)
	sort.Sort(ByTime(people))

	nowAlive := make(map[int]struct{}, n)
	collections := make([]map[int]struct{}, 0, n)

	for _, p := range people {
		if p.event == birth {
			nowAlive[p.id] = struct{}{}
			continue
		}
		if len(nowAlive) == 0 {
			delete(nowAlive, p.id)
			continue

		}
		if in(collections, nowAlive) {
			delete(nowAlive, p.id)
			continue
		}

		set := make(map[int]struct{}, len(nowAlive))
		for k := range nowAlive {
			set[k] = struct{}{}
		}

		collections = append(collections, set)
		delete(nowAlive, p.id)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, set := range collections {
		for id := range set {
			fmt.Fprint(w, id, " ")
		}
		fmt.Fprintln(w)
	}
	if len(collections) == 0 {
		fmt.Fprintln(w, 0)
	}
}

func in(collections []map[int]struct{}, arr map[int]struct{}) bool {
	for _, c := range collections {
		cnt := 0
		for k := range c {
			if _, ok := arr[k]; ok {
				cnt++
			}
		}
		if cnt == len(arr) {
			return true
		}
	}
	return false
}

func readPeople(r io.Reader, n int) []man {
	people := make([]man, 0, n*2)
	for i := 0; i < n; i++ {
		var d, m, y int
		fmt.Fscan(r, &d, &m, &y)
		bd := man{
			t:     time.Date(y+18, time.Month(m), d, 0, 0, 0, 0, time.UTC),
			id:    i + 1,
			event: birth,
		}
		maxAge := man{
			t:     time.Date(y+80, time.Month(m), d, 0, 0, 0, 0, time.UTC),
			id:    i + 1,
			event: death,
		}

		fmt.Fscan(r, &d, &m, &y)
		dd := man{
			t:     time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC),
			id:    i + 1,
			event: death,
		}
		if dd.t.Sub(bd.t) <= 0 {
			continue
		}
		if dd.t.Sub(maxAge.t) > 0 {
			dd = maxAge
		}
		people = append(people, bd, dd)
	}
	return people
}
