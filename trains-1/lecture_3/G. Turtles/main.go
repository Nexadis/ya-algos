package main

import "fmt"

func main() {
	turtles := input()
	ans := solution(turtles)
	fmt.Println(ans)
}

type turtle struct {
	before int
	after  int
}

func input() []turtle {
	var n int
	fmt.Scan(&n)
	turtles := make([]turtle, 0, n)
	for i := 0; i < n; i++ {
		t := turtle{}
		fmt.Scan(&t.before, &t.after)
		if t.before+t.after+1 != n {
			continue
		}
		if t.before < 0 || t.after < 0 {
			continue
		}
		turtles = append(turtles, t)
	}
	return turtles
}

func solution(turtles []turtle) int {
	best := 0
	pos := make(map[int]struct{}, len(turtles))
	for _, v := range turtles {
		if _, ok := pos[v.before]; ok {
			continue
		}
		pos[v.before] = struct{}{}
		best += 1

	}
	return best
}
