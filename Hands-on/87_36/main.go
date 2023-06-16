package main

import (
	"fmt"
)

func main() {
	//data structure - map
	m := map[string]int{
		"James": 42,
		"Bond":  32,
	}
	for k, v := range m {
		fmt.Printf("key is %v \t value is %v\n", k, v)
	}
}
