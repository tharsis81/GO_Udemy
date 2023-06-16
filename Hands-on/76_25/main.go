package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(400)
	fmt.Printf("The value of x is %v\n", x)
	/*
		if x <= 100 {
			fmt.Println("Less then 100\t")
		} else if x >= 101 && x <= 200 {
			fmt.Println("Between 101 and 200\t")
		} else if x >= 201 && x <= 250 {
			fmt.Println("Between 201 and 250\t")
		} else {
			fmt.Println("more then 250\t")
		}
	*/

	switch {
	case x <= 100:
		fmt.Println("Less Then 100")
	case x >= 101 && x <= 200:
		fmt.Println("Between 101 and 200\t")
	case x >= 201 && x <= 250:
		fmt.Println("Between 201 and 250\t")
	default:
		fmt.Println("more then 250\t")
	}

	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
	fmt.Println(rand.Intn(3))
}
