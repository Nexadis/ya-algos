package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	list := input(n)
	from, to := solution(list)
	fmt.Printf("%.6f %.6f\n", from, to)
}

type status byte

const (
	closer  status = 0
	further status = 1
)

type measurement struct {
	num        float64
	difference status
}

func input(n int) []measurement {
	var status string
	m := make([]measurement, n)
	fmt.Scan(&m[0].num)
	for i := 1; i < n; i++ {
		fmt.Scan(&m[i].num, &status)
		switch status {
		case "closer":
			m[i].difference = closer
		case "further":
			m[i].difference = further
		}
	}
	return m
}

func solution(m []measurement) (float64, float64) {
	minFreq := 30.0
	maxFreq := 4000.0
	for i := 1; i < len(m); i++ {
		if math.Abs(m[i].num-m[i-1].num) < math.Pow10(-6) {
			continue
		}
		if m[i].difference == closer {
			if m[i].num > m[i-1].num {
				minFreq = max((m[i].num+m[i-1].num)/2, minFreq)
			} else {
				maxFreq = min((m[i-1].num+m[i].num)/2, maxFreq)
			}
		} else {
			if m[i].num > m[i-1].num {
				maxFreq = min((m[i].num+m[i-1].num)/2, maxFreq)
			} else {
				minFreq = max((m[i-1].num+m[i].num)/2, minFreq)
			}
		}
	}
	return minFreq, maxFreq
}
