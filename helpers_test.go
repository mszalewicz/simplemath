package main

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {

	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1244, 903452, 593439128}, 4},
		{[]int{24634586, 234529626, 68395906}, 2},
		{[]int{15456, 24, 17850}, 6},
		{[]int{2, 11}, 1},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("GCD test %d", index), func(t *testing.T) {
			result := GCD(test.input[0], test.input[1], test.input[2:]...)
			if result != test.expected {
				t.Fatalf("GDC fail - expected: %d | result: %d", test.expected, result)
			}
		})
	}
}
