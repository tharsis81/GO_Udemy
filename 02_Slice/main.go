/*
+slice je samo pointer na underying array
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// si := []string{"a", "b", "c"}
	// fmt.Println(si)

	{
		fmt.Println("--------------Make -  slices---------------------------------------------------------------------------")
		/*z make pripravimo uderlying array da drži 10 vrednosti, začnemo pa s praznim arrayem. Tako naredimo če vemo koliko
		vrednosti bomo zapisali v slice. Tako je hitreje kot če ga kasneje širimo!!*/
		xi := make([]int, 0, 10)
		fmt.Println(xi)
		fmt.Println(len(xi))
		fmt.Println(cap(xi)) //slice capacity
		xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
		fmt.Println(xi)
		fmt.Println("------------")
		fmt.Println(len(xi))
		fmt.Println(cap(xi))
		fmt.Println("------------")
		xi = append(xi, 10, 11, 12, 13)
		fmt.Println(xi)
		fmt.Println(len(xi))
		fmt.Println(cap(xi))
	}

	{
		fmt.Println("--------------Multi-dimensional slices (slice of slice)--------------------------------------------")
		jb := []string{"James", "Bond", "Martini", "Chocolate"}
		jm := []string{"Jenny", "Moneypenny", "Guiness", "Wolverine Tracks"}
		fmt.Println(jb)
		fmt.Println(jm)

		xxs := [][]string{jb, jm}
		fmt.Println(xxs)

	}

	{
		fmt.Println("--------------GO internals01----------------------------------------------------------------------")
		a := []int{0, 1, 2, 3, 4, 5}
		// b := a
		//naredimo raje nov array da lahko spreminjamo elemente ločeno za vsak slice posebej
		b := make([]int, 6)
		//copy slice a into b
		copy(b, a)

		fmt.Println("a ", a)
		fmt.Println("b ", b)
		fmt.Println("--------------")

		a[0] = 7

		fmt.Println("a ", a)
		fmt.Println("b ", b)
		fmt.Println("--------------")
	}

	{
		fmt.Println("--------------GO internals02----------------------------------------------------------------------")
		n := []float64{3, 1, 4, 2}

		fmt.Println(medianOne(n))
		fmt.Println("after medianOne", n)

		y := []float64{3, 1, 4, 2}
		fmt.Println(medianTwo(y))
		fmt.Println("after medianTwo", y)

	}
}
func medianOne(x []float64) float64 {
	sort.Float64s(x)
	i := len(x) / 2
	if len(x)%2 == 1 {
		return x[i/2]
	}
	return (x[i-1] + x[i]) / 2
}

func medianTwo(x []float64) float64 {
	// allocate a new underlying array
	n := make([]float64, len(x))
	copy(n, x)

	sort.Float64s(n)
	i := len(n) / 2
	if len(n)%2 == 1 {
		return n[i/2]
		// when you divide
		// you get the whole number quotient
		// without the remainder modulus
		// https://go.dev/play/p/2r5WrelUEh7
	}
	return (n[i-1] + n[i]) / 2
}
