package main

import (
	"fmt"
)

func main() {

	fmt.Println("-----------------------------------------------------------")
	//data structure - map
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("key is %v \t value is %v\n", k, v)
	}

	fmt.Println("-----------------------------------------------------------")

	age := m["James"]
	fmt.Println("The age of Bond", age)

	age = m["Q"]
	fmt.Println("The age of Q", age)
	//comma ok idiom
	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q", v)
	}

}
