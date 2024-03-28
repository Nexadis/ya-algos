package main

import "fmt"

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() int {
	digits := make(map[int]struct{}, 3)
	nums := make(map[int]struct{}, 3)
	for i := 0; i < 3; i++ {
		var x int
		fmt.Scan(&x)
		digits[x] = struct{}{}
	}
	var n, ans int
	fmt.Scan(&n)

	for n > 0 {
		d := n % 10
		nums[d] = struct{}{}
		n = n / 10
	}
	for k := range nums {
		if _, ok := digits[k]; !ok {
			ans += 1
		}
	}
	return ans
}
