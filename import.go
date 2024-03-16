package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Nichola")
	fmt.Println(result)

	fmt.Println(helper.Application)
	fmt.Println(helper.Contoh("Nichola"))

}
