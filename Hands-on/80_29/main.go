package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// for i := 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	for i := 0; i < 100; i++ {

		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("Iteration %v \t The value of x is %v and the value of y is %v\n", i, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Println("Both are less than 4")
		case x > 6 && y > 6:
			fmt.Println("Both are greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("X is between 4 and 6")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("none were met")
		}

	}
}
