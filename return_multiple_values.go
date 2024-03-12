package main

import "fmt"

func getFullName() (string, string) {
	return "Hello", "Nichola"
}

func main() {
	sayHello, name := getFullName()

	fmt.Println(sayHello, name)
}
