package main

import (
	"fmt"
)

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	result := sumAll(10, 10, 10, 10)

	fmt.Println(result)

	numbers := []int{10, 10, 10, 10, 100}

	fmt.Println(sumAll(numbers...))
}
