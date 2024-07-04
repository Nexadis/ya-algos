package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type group struct {
	id       int
	capacity int
	isBusy   bool
}

func main() {
	rb := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(rb, &n, &m)

	students := make([]group, n)
	for i := range students {
		var c int
		fmt.Fscan(rb, &c)
		students[i] = group{id: i + 1, capacity: c + 1}
	}

	rooms := make([]group, m)
	for i := range rooms {
		var c int
		fmt.Fscan(rb, &c)
		rooms[i] = group{id: i + 1, capacity: c}
	}

	sort.Slice(students, ByCap(students))
	sort.Slice(rooms, ByCap(rooms))

	ans := make([]int, n+1)
	good := 0
	for _, s := range students {
		for j := 0; j < len(rooms); j++ {
			if rooms[j].capacity >= s.capacity && !rooms[j].isBusy {
				good++
				ans[s.id] = rooms[j].id
				rooms[j].isBusy = true
				break
			}
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	fmt.Fprintln(w, good)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(w, ans[i])
	}
}

func ByCap(grps []group) (less func(int, int) bool) {
	return func(i, j int) bool {
		return grps[i].capacity > grps[j].capacity
	}
}
