package main

import "fmt"

func getGoodBye(name string) string {

	return "Hello " + name
}

func main() {
	contoh := getGoodBye

	fmt.Println(contoh("Nichola"))
}
