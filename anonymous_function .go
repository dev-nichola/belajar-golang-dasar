package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You Are blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "GUGUK"
	}
	registerUser("GUGUK", blacklist)
	registerUser("Nichola", func(name string) bool {
		return name == "ANJING"
	})
}
