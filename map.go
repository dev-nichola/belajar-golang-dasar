package main

import "fmt"

func main() {

	// Cara 1
	var person1 map[string]string = map[string]string{}

	person1["name"] = "Nichola"
	person1["address"] = "Blitar"

	// Cara 2
	person2 := map[string]string{
		"name":    "Nichola",
		"address": "Blitar",
	}

	fmt.Println(person1["name"])
	fmt.Println(person1["address"])

	fmt.Println(person2["name"])
	fmt.Println(person2["address"])

	book := make(map[string]string)
	book["title"] = "Buku Matematika"
	book["author"] = "Nichola"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)

}
