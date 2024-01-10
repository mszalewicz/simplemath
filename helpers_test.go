package main

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {

	functionName := "GCD"

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
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {
			result := GCD(test.input[0], test.input[1], test.input[2:]...)
			if result != test.expected {
				t.Fatalf("%s fail - expected: %d | result: %d", functionName, test.expected, result)
			}
		})
	}
}

func TestLCM(t *testing.T) {

	functionName := "LCM"

	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{12, 15, 75}, 300},
		{[]int{14, 19, 122}, 16226},
		{[]int{7, 5}, 35},
	}

	for index, test := range tests {
		t.Run(fmt.Sprintf("%s test %d", functionName, index), func(t *testing.T) {
			result := LCM(test.input[0], test.input[1], test.input[2:]...)
			if result != test.expected {
				t.Fatalf("%s fail - expected: %d | result: %d", functionName, test.expected, result)
			}
		})
	}
}
