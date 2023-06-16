package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("This is an outer loop")
		for i := 0; i < 5; i++ {
			fmt.Println("This is an inner loop")
		}
	}
}
