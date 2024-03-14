package main

import "fmt"

func endApp() {
	fmt.Println("End App")

	// contoh yang benar
	message := recover()
	fmt.Println("terjadi error", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups Error")
	}

	// cara yang salah
	// message := recover()

	// fmt.Println("terjadi error", message)
}

func main() {
	runApp(true)
}
