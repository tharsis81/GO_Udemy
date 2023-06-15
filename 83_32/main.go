package main

import (
	"fmt"
)

func main() {

	y := 1

	for {
		fmt.Printf("x is %v\n", y)
		if y > 20 {
			break
		}
		y++
	}
}
