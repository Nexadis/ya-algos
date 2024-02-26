package main

import (
	"fmt"
	"strings"
)

func main() {
	var wa string
	fmt.Scanln(&wa)
	simpleTelephone := func(tel string) string {
		tel = strings.ReplaceAll(tel, "-", "")
		tel = strings.ReplaceAll(tel, "(", "")
		tel = strings.ReplaceAll(tel, ")", "")
		tel = strings.ReplaceAll(tel, "+7", "8")
		if len(tel) == 11 {
			return tel
		}
		if len(tel) == 7 {
			return "8495" + tel

		}
		return ""
	}
	wa = simpleTelephone(wa)
	var cur string
	for i := 0; i < 3; i++ {
		n, _ := fmt.Scanln(&cur)
		if n == 0 {
			break
		}
		cur = simpleTelephone(cur)
		if strings.Compare(wa, cur) != 0 {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	}
}
