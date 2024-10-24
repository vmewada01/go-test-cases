package slices

import (
	"fmt"
	"slices"
	"testing"
)

func TestSlice(t *testing.T) {
	var numbers []int
	numbers = append(numbers, 1, 2, 3, 4, 5)

	fmt.Println("Original slice:", numbers)

	numbers = append(numbers[:2], numbers[3:]...)

	if slices.Index(numbers, 6) >= 0 {
		t.Fatal("Expected error for index 6")
	}
	fmt.Println("Slicing the slice from index 1 to 3:")

	fmt.Println("After removing 3rd and 4th elements:", numbers)
	fmt.Println("Length of the slice:", len(numbers))
	fmt.Println("Capacity of the slice:", cap(numbers))
	fmt.Println("Is the slice empty?", numbers == nil)
	fmt.Println("First element:", numbers[0])
	fmt.Println("Last element:", numbers[len(numbers)-1])
	fmt.Println("Index of 4:", slices.Index(numbers, 4))
	fmt.Println("Index of 6:", slices.Index(numbers, 6))

}
