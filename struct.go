package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
}

func main() {
	var nichola Customer
	nichola.Name = "Nichola"
	nichola.Address = "Blitar"
	nichola.Age = 20

	fmt.Println(nichola)

	joko := Customer{
		Name:    "Joko",
		Address: "Malang",
		Age:     20,
	}

	fmt.Println(joko.Name)

	budi := Customer{"Budi", "Surabaya", 40}

	fmt.Println(budi)
}
