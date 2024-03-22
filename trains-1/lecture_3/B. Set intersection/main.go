package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strings"
)

func main() {
	list1, list2 := input()
	ans1 := solution(list1, list2)
	for _, v := range ans1 {
		fmt.Print(v, " ")
	}
}

func input() ([]*big.Int, []*big.Int) {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	list1 := readLine(s)
	list2 := readLine(s)
	return list1, list2
}

func solution(l1, l2 []*big.Int) []*big.Int {
	set1 := make(map[string]int, len(l1))
	set2 := make(map[string]int, len(l2))
	for _, v := range l1 {
		set1[v.String()] += 1
	}
	for _, v := range l2 {
		set2[v.String()] += 1
	}
	ans := make(bigSort, 0, 10)
	for k := range set1 {
		_, ok := set2[k]
		if ok {
			val := new(big.Int)
			val.SetString(k, 10)
			ans = append(ans, val)
		}
	}

	sort.Sort(ans)

	return ans
}

type bigSort []*big.Int

func (a bigSort) Len() int           { return len(a) }
func (a bigSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bigSort) Less(i, j int) bool { return a[i].Cmp(a[j]) == -1 }

func readLine(s *bufio.Scanner) []*big.Int {
	s.Scan()
	line := strings.TrimSpace(s.Text())
	splited := strings.Split(line, " ")
	list := make([]*big.Int, 0, len(splited))
	for _, v := range splited {
		if strings.TrimSpace(v) == "" {
			continue
		}
		i := new(big.Int)
		_, ok := i.SetString(v, 10)
		if !ok {
			continue
		}
		list = append(list, i)
	}
	return list
}
