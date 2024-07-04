package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	cnt := make([]int, 2)
	for _, b := range s {
		if b == '(' {
			cnt[0]++
			continue
		}
		cnt[1]++
		if cnt[0] < cnt[1] {
			break
		}
	}
	if cnt[0] == cnt[1] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
