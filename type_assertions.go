package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()

	var resultString string = result.(string)
	fmt.Println(resultString)

	// var resultInt int = result.(int) // panic
	// fmt.Println(resultInt)

	// Best Practice

	data := random()

	switch value := data.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unkown")
	}
}
