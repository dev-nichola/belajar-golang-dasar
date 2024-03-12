package main

import "fmt"

func main() {
	name := "Nichola"

	switch name {
	case "Budi":
		fmt.Println("Hello Budi")
	case "Joko":
		fmt.Println("Hello Joko")
	case "Nichola":
		fmt.Println("Hello Nichola")
	default:
		fmt.Println("Bukan Siapa Siapa")
	}

	lenght := len(name)
	switch {
	case lenght > 10:
		fmt.Println("Nama Terlalu Panjang")
	case lenght > 5:
		fmt.Println("Nama Anda Pas Dan Cukup")
	default:
		fmt.Println("Anda Punya Nama")
	}
}
