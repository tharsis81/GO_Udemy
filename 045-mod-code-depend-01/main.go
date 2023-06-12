package main

import (
	"fmt"

	puppy "github.com/tharsis81/puppy_tmp"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Barks())
}
