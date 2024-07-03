package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type topic struct {
	name   string
	parent int
	id     int
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()

	n, _ := strconv.Atoi(s.Text())
	topics := make(map[int]topic, n)
	i := 0

	for {
		i++
		s.Scan()
		line := s.Text()
		if line == "" {
			break
		}
		num, _ := strconv.Atoi(line)
		if num == 0 {
			s.Scan()
			name := s.Text()
			s.Scan()
			topics[i] = topic{name: name, id: i, parent: 0}
		} else {
			s.Scan()
			topics[i] = topic{id: i, parent: num}
		}
	}

	counters := make([]int, n+1)
	maxCnt := 0
	for _, t := range topics {
		t := t
		for t.parent != 0 {
			counters[t.id]++
			t = topics[t.parent]
		}
		counters[t.id]++
		maxCnt = max(maxCnt, counters[t.id])
	}
	for i, cnt := range counters {
		if cnt == maxCnt {
			fmt.Println(topics[i].name)
			break
		}
	}
}
