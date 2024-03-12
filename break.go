package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println("Perulangan Ke", i)

		if i == 5 {
			break
		}
	}

	fmt.Println("Kode Selesai Di Eksekusi")

}
