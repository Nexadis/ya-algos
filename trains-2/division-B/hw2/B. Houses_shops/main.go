package main

import "fmt"

func main() {
	city := make([]int, 10)

	nextShop := -1
	for i := 0; i < 10; i++ {
		var x int
		fmt.Scan(&x)
		if nextShop == -1 && x == 2 {
			nextShop = i
		}
		city[i] = x
	}
	prevShop := -100
	maxDiff := 0
	for i, h := range city {
		switch h {
		case 2:
			prevShop = i
			j := i + 1
			for ; j < len(city) && city[j] != 2; j++ {
			}
			nextShop = j
			if j == len(city) {
				nextShop = 100
			}
			continue
		case 0:
		case 1:
			diff := min(i-prevShop, nextShop-i)
			if diff > maxDiff {
				maxDiff = diff
			}
		}
	}
	fmt.Println(maxDiff)
}
