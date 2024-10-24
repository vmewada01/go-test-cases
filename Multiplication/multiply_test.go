package multiplication

import "testing"

func TestMultiplyTwoNumber(t *testing.T) {
	// Test cases
	testCases := []struct {
		num1, num2     int
		expectedResult int
	}{
		{2, 3, 6},
		{10, 5, 50},
		{-4, 6, -24},
		{0, 10, 0},
	}

	// Run the test cases
	for _, testCase := range testCases {
		result := MultiplyTwoNumber(testCase.num1, testCase.num2)
		if result != testCase.expectedResult {
			t.Errorf("Expected %d, but got %d for %d * %d", testCase.expectedResult, result, testCase.num1, testCase.num2)
		}
	}
	t.Log("All test cases passed")
}
