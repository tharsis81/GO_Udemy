package main

import "fmt"

var y int

func main() {
	fmt.Println(y)

	z := 42
	fmt.Println(z)

	a, b := 43, "HAppines"
	fmt.Println(a, b)

	var s float32 = 44.44
	fmt.Printf("%v is of this type %T\n", s, s)

	e, f, _ := 45, 46, 47
	fmt.Println(e, f)
}
