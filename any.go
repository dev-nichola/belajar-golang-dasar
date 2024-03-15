package main

import "fmt"

func Ups() any {
	return "Ups Salah"
}

func main() {
	var kosong any = Ups()

	fmt.Println(kosong)
}
