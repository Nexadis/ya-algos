package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	x := search(func(x float64) bool {
		y := float64(a)*x*x*x + float64(b)*x*x + float64(c)*x + float64(d)
		if a < 0 {
			return y < 0
		}
		return y > 0
	})
	fmt.Println(x)
}

func search(isLeft func(m float64) bool) float64 {
	l := float64(-2 << 20)
	r := float64(2 << 20)
	for r-l > 0.000_000_1 {
		m := (r + l) / 2
		if isLeft(m) {
			r = m
		} else {
			l = m
		}
	}
	return l
}
