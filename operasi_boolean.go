package main

import "fmt"

func main() {
	nilaiAkhir := 90
	nilaiAbsensi := 80
	nilaiLulus := 70

	var nilaiLulusAkhir bool = nilaiAkhir > nilaiLulus
	var nilaiAbsensiAkhir bool = nilaiAbsensi >= 80

	var lulus bool = nilaiAbsensiAkhir && nilaiLulusAkhir

	fmt.Println(lulus)

}
