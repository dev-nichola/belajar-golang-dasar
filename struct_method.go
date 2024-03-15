package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
	// sayHello() seperti ini logikanya
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My Name Is", customer.Name)
}

func main() {
	joko := Customer{Name: "Joko"}

	joko.sayHello()
}
