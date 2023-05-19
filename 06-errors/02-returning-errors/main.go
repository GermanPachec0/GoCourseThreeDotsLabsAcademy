package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := Divide(100, 50)
	fmt.Println("Result:", result, "Error:", err)
}

func Divide(num1, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, errors.New("Argument Cannot be 0")
	}
	return num1 / num2, nil
}
