package main

import "fmt"

func main() {
	a := 10
	b := 10
	d := 5
	e := 2
	f := 2

	c := a + b - d*e/f

	fmt.Println(c)

	// Operasi Penugasa
	i := 10

	i += 10
	fmt.Println(i)

	i += 10
	fmt.Println(i)

	// Operasi Unary
	x := 1
	x++
	x++
	fmt.Println(x)
	x--
	fmt.Println(x)
}
