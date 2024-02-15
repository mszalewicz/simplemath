package simplemath

import (
	"math"
	"slices"
)

// Euclidian algorithm for finding greatest common divisor for a set of integers
// https://en.wikipedia.org/wiki/Euclidean_algorithm
func GCD(x, y int, numbers ...int) int {
	for {
		if y == 0 {
			break
		}

		t := y
		y = x % y
		x = t
	}

	for i := 0; i < len(numbers); i++ {
		x = GCD(x, numbers[i])
	}

	return x
}

// Implementation of least common multiple for a set of integers
// https://en.wikipedia.org/wiki/Least_common_multiple
func LCM(x, y int, numbers ...int) int {
	result := x * y / GCD(x, y)

	for i := 0; i < len(numbers); i++ {
		result = LCM(result, numbers[i])
	}

	return result
}

// Return distance in 1D space
func Distance(x, y float64) float64 {
	if (x < 0 && y < 0) || (x > 0 && y > 0) {
		return math.Max(math.Abs(x), math.Abs(y)) - math.Min(math.Abs(x), math.Abs(y))
	} else {
		return math.Abs(x) + math.Abs(y)
	}
}

func ProperDivisors(number int) []int {
	if number <= 1 {
		return nil
	}

	properDivisors := []int{}

	for i := 1; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			properDivisors = append(properDivisors, i)
			if number != i*i && number/i != number {
				properDivisors = append(properDivisors, number/i)
			}
		}
	}

	slices.Sort(properDivisors)

	return properDivisors
}
