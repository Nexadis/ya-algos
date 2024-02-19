package main

import "fmt"

func main() {
	var troom, tcond int
	var mode string
	fmt.Scan(&troom, &tcond, &mode)
	switch mode {
	case "heat":
		if troom > tcond {
			fmt.Println(troom)
		} else {
			fmt.Println(tcond)
		}
	case "freeze":
		if troom < tcond {
			fmt.Println(troom)
		} else {
			fmt.Println(tcond)
		}
	case "auto":
		fmt.Println(tcond)
	case "fan":
		fmt.Println(troom)
	}
}
