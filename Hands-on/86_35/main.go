package main

import (
	"fmt"
)

func main() {
	//data structure - a slice of ints
	xi := []int{42, 43, 44, 45, 46}
	for i, v := range xi {
		fmt.Println(i, v)
	}
}
