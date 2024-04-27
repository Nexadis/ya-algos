package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	m, t1, t2 := solution(os.Stdin)
	fmt.Println(m, t1, t2)
}

const (
	comeIn = iota
	cameOut
)

type event struct {
	t         int
	id        int
	typeEvent int
}

type ByTimeEnter []event

func (a ByTimeEnter) Len() int      { return len(a) }
func (a ByTimeEnter) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTimeEnter) Less(i, j int) bool {
	if a[i].t == a[j].t {
		return a[i].typeEvent < a[j].typeEvent
	}
	return a[i].t < a[j].t
}

func solution(r io.Reader) (int, int, int) {
	rb := bufio.NewReader(r)
	var n int
	fmt.Fscan(rb, &n)
	events := make([]event, 0, n*2)
	for i := 0; i < n; i++ {
		from := event{id: i, typeEvent: comeIn}
		to := event{id: i, typeEvent: cameOut}
		fmt.Fscan(rb, &from.t, &to.t)
		if to.t-from.t < 5 {
			continue
		}
		to.t -= 5
		events = append(events, from, to)
	}
	sort.Sort(ByTimeEnter(events))
	return adsChoose(events)
}

func adsChoose(events []event) (int, int, int) {
	firstAd := make(map[int]struct{}, len(events))
	maxClients := 0
	firstBest := 0
	secondBest := 0
	if len(events) == 0 {
		return 0, 10, 20
	}
	if len(events) == 2 {
		return 1, events[0].t, events[0].t + 10
	}
	for i, e := range events {
		if e.typeEvent == comeIn {
			firstAd[e.id] = struct{}{}
			if len(firstAd) > maxClients {
				maxClients = len(firstAd)
				firstBest = e.t
				secondBest = e.t + 5
			}
		}
		secondAdCnt := 0
		for j := i + 1; j < len(events); j++ {
			e2 := events[j]

			if _, ok := firstAd[e2.id]; e2.typeEvent == comeIn && !ok {
				secondAdCnt++
			}
			if e2.t-5 >= e.t && len(firstAd)+secondAdCnt > maxClients {
				maxClients = len(firstAd) + secondAdCnt
				firstBest = e.t
				secondBest = e2.t
			}
			if _, ok := firstAd[e2.id]; e2.typeEvent == cameOut && !ok {
				secondAdCnt--
			}
		}
		if e.typeEvent == cameOut {
			delete(firstAd, e.id)
		}
	}
	return maxClients, firstBest, secondBest
}
