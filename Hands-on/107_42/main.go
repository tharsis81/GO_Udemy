package main

import "fmt"

func main() {
	//create an array
	x := [5]int{}
	//assign values to array
	for i := 0; i < 5; i++ {
		x[i] = i
	}

	//print out
	//omitamo i da ga ne rabimo uporabit
	for _, v := range x {
		fmt.Printf("%v %T\n", v, v)
	}

}
