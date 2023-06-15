package main

import (
	"fmt"
	"math/rand"
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

	age1 := m["James"]
	fmt.Println("The age of Bond", age1)

	age2 := m["Q"]
	fmt.Println("The age of Q", age2)
	//comma ok idiom
	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q", v)
	}

	fmt.Println("-----------------------------------------------------------")

	c := 1
	for i := 0; i < 100; i++ {
		//statement-statement idiom
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Iteration %v \t Total count %v \t X is %v\n", i, c, x)
			c++
		}
	}
}
