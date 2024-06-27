package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	arrive = iota
	departure
)

type track struct {
	t      int
	city   int
	action int
}

type ByTime []track

func (a ByTime) Len() int      { return len(a) }
func (a ByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool {
	if a[i].t == a[j].t {
		return a[i].action < a[j].action
	}
	return a[i].t < a[j].t
}

func main() {
	r := bufio.NewReader(os.Stdin)
	ans := solution(r)
	fmt.Println(ans)
}

func solution(r io.Reader) int {
	var n, m int
	fmt.Fscan(r, &n, &m)

	tracks := make([]track, 0, m*4)

	cityStat := make(map[int]int, n)
	for i := 0; i < m; i++ {
		var fromCity, toCity int
		var from, to string
		fmt.Fscan(r, &fromCity, &from, &toCity, &to)
		tfrom, tto := parse(from), parse(to)
		if tto < tfrom {
			tto += 24 * 60
			cityStat[toCity]++
		} else {
			tracks = append(tracks,
				track{t: tto + 24*60, city: toCity, action: arrive},
			)
		}
		tracks = append(tracks,
			track{t: tfrom, city: fromCity, action: departure},
			track{t: tto, city: toCity, action: arrive},
			track{t: tfrom + 24*60, city: fromCity, action: departure},
		)
	}

	sort.Sort(ByTime(tracks))
	busParked := make(map[int]int, n)
	allNeedBus := 0

	for _, t := range tracks {
		if t.action == departure {
			cityStat[t.city]--
			if busParked[t.city] <= 0 {
				allNeedBus++
			} else {
				busParked[t.city]--
			}
		}
		if t.action == arrive {
			cityStat[t.city]++
			busParked[t.city]++
		}
	}

	for _, c := range cityStat {
		if c != 0 {
			return -1
		}
	}

	return allNeedBus
}

func parse(s string) int {
	t := strings.Split(s, ":")
	h, _ := strconv.Atoi(t[0])
	m, _ := strconv.Atoi(t[1])
	return h*60 + m
}
