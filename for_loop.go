package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan Ke = ", counter)
	}

	fmt.Println("Selesai")

	names := []string{"Nichola", "Saputra"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("INDEX = ", index, "NAME = ", name)
	}

	// Tidak Butuh Indexnya
	for _, name := range names {
		fmt.Println("NAME = ", name)
	}

}
