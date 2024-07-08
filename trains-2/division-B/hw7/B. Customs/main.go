package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	finish = -1
	begin  = 1
)

type load struct {
	when int
	t    int
}

func main() {
	rb := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(rb, &n)

	goods := make([]load, 0, n*2)
	for i := 0; i < n; i++ {
		var start, cont int
		fmt.Fscan(rb, &start, &cont)
		goods = append(goods,
			load{when: start, t: begin},
			load{when: start + cont, t: finish},
		)
	}

	sort.Slice(goods, func(i, j int) bool {
		if goods[i].when == goods[j].when {
			return goods[i].t < goods[j].t
		}
		return goods[i].when < goods[j].when
	})

	cnt := 0
	maxCnt := 0
	for _, g := range goods {
		if g.t == begin {
			cnt++
		}
		if g.t == finish {
			cnt--
		}
		maxCnt = max(maxCnt, cnt)
	}
	fmt.Println(maxCnt)
}
