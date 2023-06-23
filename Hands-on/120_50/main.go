package main

import (
	"fmt"
)

func main() {
	m := make(map[string][]string)
	m["bond_james"] = []string{"test", "test6"}
	m["anze_pirc"] = []string{"test2", "test5"}
	m["mišek"] = []string{"test3"}
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Println(i, v2)
		}
	}
}
