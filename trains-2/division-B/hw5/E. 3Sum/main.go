package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type num struct {
	id  int
	val int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var s int
	fmt.Fscan(r, &s)

	as := readSeq(r)
	bs := readSeq(r)
	cs := readSeq(r)
	ans := position{}
	found := false
	for _, a := range as {
		cpos := len(cs) - 1
		for _, b := range bs {
			for cpos > 0 && a.val+b.val+cs[cpos].val > s {
				cpos--
			}
			if a.val+b.val+cs[cpos].val == s && (!found || lessPos(position{a.id, b.id, cs[cpos].id}, ans)) {
				ans = position{a.id, b.id, cs[cpos].id}
				found = true
			}
		}
	}
	if !found {
		fmt.Println(-1)
	} else {
		fmt.Println(ans.ai, ans.bi, ans.ci)
	}
}

func readSeq(r *bufio.Reader) []pair {
	var n int
	fmt.Fscan(r, &n)
	nums := make([]pair, n)
	for i := range nums {
		var v int
		fmt.Fscan(r, &v)
		nums[i] = pair{id: i, val: v}
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i].val == nums[j].val {
			return nums[i].id > nums[j].id
		}
		return nums[i].val < nums[j].val
	})
	return nums
}

type pair struct {
	id  int
	val int
}

type position struct {
	ai int
	bi int
	ci int
}

func lessPos(p1, p2 position) bool {
	if p1.ai == p2.ai {
		if p1.bi == p2.bi {
			return p1.ci < p2.ci
		}
		return p1.bi < p2.bi
	}
	return p1.ai < p2.ai
}
