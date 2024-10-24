package slices

import "fmt"

func Slice() {
	// Declare and initialize a slice of integers
	slice := []int{1, 2, 3, 4, 5}

	// Print the length and capacity of the slice
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))

	// Append a new element to the slice
	slice = append(slice, 6)

	// Print the updated length and capacity of the slice
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))

	// Print the slice
	fmt.Println(slice)

	// Slice a portion of the slice
	newSlice := slice[1:3]

	// Print the new slice
	fmt.Println(newSlice)

	// Modify the new slice
	newSlice[0] = 100

	// Print the modified slice and the original slice
	fmt.Println(newSlice)
	fmt.Println(slice)

}
