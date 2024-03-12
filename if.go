package main

import "fmt"

func main() {
	name := "Nichola Saputra"

	if name == "Nichola" {
		fmt.Println("Hello Nichola")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Bukan Keduanya")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}

}
