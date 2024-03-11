package main

import "fmt"

func main() {
	type NoKTP string

	var ktp NoKTP = "111111111111"
	fmt.Println(ktp)
	fmt.Println()
}
