package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "Validation Error"}
	}

	if id != "nichola" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}
func main() {
	err := saveData("nichol", nil)

	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error :", validationErr.Message)
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("not found error : ", notFoundError.Message)
		} else {
			fmt.Println("unkown error : ", err.Error())
		}
	}
}
