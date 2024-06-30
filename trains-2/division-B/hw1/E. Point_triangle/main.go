package main

import "fmt"

func main() {
	var d, xx, xy int
	fmt.Scan(&d, &xx, &xy)

	switch {
	case (xx >= 0 && xx <= d) && (xy >= 0 && xy <= d) && xx+xy <= d:
		fmt.Println(0)
	case xx <= d/2 && xy <= d/2:
		fmt.Println(1)
	case xx >= xy:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}
}
