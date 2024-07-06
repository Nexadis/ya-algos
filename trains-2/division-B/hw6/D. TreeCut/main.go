package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, k, b, m, x int64
	fmt.Scan(&a, &k, &b, &m, &x)
	l := int64(0)
	r := x/a*2 + 1

	need := big.NewInt(x)
	d := search(l, r, func(d int64) bool {
		biga := big.NewInt(a)
		ha := big.NewInt(d - d/k)
		bigb := big.NewInt(b)
		hb := big.NewInt(d - d/m)
		biga = biga.Mul(biga, ha)
		bigb = bigb.Mul(bigb, hb)
		cut := biga.Add(biga, bigb)
		return cut.Cmp(need) < 0
	})
	fmt.Println(d)
}

func search(l, r int64, f func(m int64) bool) int64 {
	for l < r {
		m := (l + r) / 2
		if f(m) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
