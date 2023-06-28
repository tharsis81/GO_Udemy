package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

//embedded struct
type secretAgent struct {
	person
	ltk   bool
	first string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   32,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.age, p1.last, p1.age)

	fmt.Printf("%T \t %#v\n", p1, p1)

	fmt.Println("--------------------------Emedded struct----------------------------------------")

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "Ethan Hawk",
		ltk:   true,
	}

	fmt.Println(sa1.first, sa1.last, sa1.age)
	fmt.Println(sa1.first, sa1.person.first)

	fmt.Println("--------------------------Anonymous struct-----------------------------------------")

	//check GO code repo!!!

	fmt.Println("--------------------------Composition-----------------------------------------")
	//read about it!!!
}
