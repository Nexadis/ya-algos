package main

import "fmt"

func main() {
	var r, i, c int8
	ans := int8(0)
	fmt.Scan(&r, &i, &c)
	switch i {
	case 0:
		if r != 0 {
			ans = 3
		} else {
			ans = c
		}
	case 1:
		ans = c
	case 4:
		if r != 0 {
			ans = 3
		} else {
			ans = 4
		}
	case 6:
		ans = 0
	case 7:
		ans = 1
	default:
		ans = i
	}
	fmt.Println(ans)
}
