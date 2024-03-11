package main

import (
	"fmt"
)

func main() {
	var nilai32 int32 = 32178000
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Nichola Saputra"
	var e = name[8]
	var eString = string(e)

	fmt.Println(eString)
}
