package main

import (
	"fmt"
	addition "test-cases/Addition"
)

func main() {
	fmt.Println("hello world")
	// Example usage of AdditionOfTwoNumber function
	num1 := 5
	num2 := 7
	result := addition.AdditionOfTwoNumber(num1, num2)
	fmt.Printf("The result of adding %d and %d is: %d\n", num1, num2, result)

	// You can also test other values
	result = addition.AdditionOfTwoNumber(-10, 15)
	fmt.Printf("The result of adding %d and %d is: %d\n", -10, 15, result)

}
