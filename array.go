package main

import "fmt"

func main() {
	var names [2]string

	names[0] = "Nichola"
	names[1] = "Saputra"

	fmt.Println(names[0])
	fmt.Println(names[1])

	// Error
	// names[2] = "Laaaa"

	var values = [...]int{
		90, 80, 70}

	fmt.Println(values)
	fmt.Println(len(values))
}
