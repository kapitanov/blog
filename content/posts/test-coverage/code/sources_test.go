package main

import (
	"fmt"
	"math"
	"testing"
)

func TestCalcDiscount(t *testing.T) {
	type TestCase struct {
		UserAge      int
		IsUserActive bool
		Expected     float64
	}

	testCases := []TestCase{
		{UserAge: 17, IsUserActive: false, Expected: 0.72}, // Test Case 1 - user age is < 18 and user is inactive
		{UserAge: 30, IsUserActive: true, Expected: 1.0},   // Test Case 2 - user age is ≥ 18 and user is active
		{UserAge: 17, IsUserActive: true, Expected: 0.9},   // Test Case 3 - user age is < 18 and user is active
		{UserAge: 25, IsUserActive: false, Expected: 0.8},  // Test Case 4 - user age is ≥ 18 and user is inactive
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("%d-%t", testCase.UserAge, testCase.IsUserActive), func(t *testing.T) {
			actual := CalcDiscount(testCase.UserAge, testCase.IsUserActive)

			delta := math.Abs(actual - testCase.Expected)
			const epsilon = 0.00001
			if delta > epsilon {
				t.Errorf("Expected %v, got %v", testCase.Expected, actual)
			}
		})
	}
}
