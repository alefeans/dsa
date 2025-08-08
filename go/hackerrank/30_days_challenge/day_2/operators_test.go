package main

import "testing"

const errorMsg = "expected value %v, found %v"

func TestSolve(t *testing.T) {
	t.Parallel()
	tests := []struct {
		mealCost   float64
		tipPercent int32
		taxPercent int32
		expected   float64
	}{
		{12.00, 20, 8, 15},
		{15.50, 15, 10, 19},
	}

	for _, test := range tests {
		t.Logf("Data test: %v", test)
		value := solve(test.mealCost, test.tipPercent, test.taxPercent)
		if value != test.expected {
			t.Errorf(errorMsg, test.expected, value)
		}
	}
}
