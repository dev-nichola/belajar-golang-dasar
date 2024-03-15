package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	alamat1 := Address{"Blitar", "Jawa Timur", "Indonesia"}

	// alamat2 := &alamat1 // Pointer
	var alamat2 *Address = &alamat1

	alamat2.City = "Malang" // dirubah menjadi malang

	fmt.Println(alamat1) // ikut berubah menjadi malang
	fmt.Println(alamat2) // berubah menjadi malang

	*alamat2 = Address{"Jakarta", "DKJ", "Indonesia"}

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}
