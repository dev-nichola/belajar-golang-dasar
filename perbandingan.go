package main

import "fmt"

func main() {
	name1 := "Nichola"
	name2 := "Saputra"

	result1 := name1 != name2

	fmt.Println(result1)

	value1 := 100
	value2 := 99

	result2 := value1 < value2
	fmt.Println(result2)
}
