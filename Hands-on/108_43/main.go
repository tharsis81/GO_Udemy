package main

import "fmt"

func main() {
	xs := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range xs {
		fmt.Printf("index %v - Value %v - Type %T\n", i, v, v)
	}
}
