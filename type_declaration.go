package main

import "fmt"

func main() {
	type NoKTP string

	var ktp NoKTP = "111111111111"

	fmt.Println(ktp)

	// Konversi Ke Tipe Data NoKTP
	var contoh string = "222222222"
	fmt.Println(NoKTP(contoh))

	var konversi NoKTP = NoKTP(contoh)
	fmt.Println(konversi)
}
