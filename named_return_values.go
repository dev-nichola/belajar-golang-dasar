package main

import "fmt"

func getCompeleteName() (firstName string, lastName string) {
	firstName = "Nichola"
	lastName = "Saputra"

	return firstName, lastName
}

func main() {
	a, _ := getCompeleteName()

	fmt.Println(a)
}
