package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Blitar", "Jawa Timur", "Indonesia"}

	address2 := address1 // Copy Value

	address2.City = "Malang"
	fmt.Println(address1)
	fmt.Println(address2)

	alamat1 := Address{"Blitar", "Jawa Timur", "Indonesia"}

	// alamat2 := &alamat1 // Pointer
	var alamat2 *Address = &alamat1

	alamat2.City = "Malang" // dirubah menjadi malang

	fmt.Println(alamat1) // ikut berubah menjadi malang
	fmt.Println(alamat2) // berubah menjadi malang

}
