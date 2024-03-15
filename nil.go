package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	result := NewMap("")

	if result == nil {
		fmt.Println("Data Masih Kosong")
	} else {
		fmt.Println(result["name"])
	}
}
