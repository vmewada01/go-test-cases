package addition

import (
	"testing"
)

func TestAdditionOfTwoNumber(t *testing.T) {
	// t.Log()
	// t.Fail()
	// t.FailNow()
	// t.Error()
	// t.Fatal()

	// Test input arrays
	num1 := []int{2, 2, -5, 9, 34, -52}
	num2 := []int{3, 4, -7, 2, 42, 7}
	expectedResults := []int{5, 6, -12, 11, 76, -45}

	// Iterate over test cases
	for i := 0; i < len(num1); i++ {
		result := AdditionOfTwoNumber(num1[i], num2[i])
		if result != expectedResults[i] {
			t.Errorf("Test failed for inputs %d and %d: expected %d, got %d", num1[i], num2[i], expectedResults[i], result)
		} else {
			t.Logf("Test passed for inputs %d and %d: expected %d, got %d", num1[i], num2[i], expectedResults[i], result)
		}
	}

	// Single test case
	want := 8
	got := AdditionOfTwoNumber(3, 5)
	if got != want {
		t.Errorf("Test failed! want: '%d', got: '%d'", want, got)
	} else {
		t.Logf("Test passed! want: '%d', got: '%d'", want, got)
	}
}
