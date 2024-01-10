package main

// Euclidian algorithm for finding greatest common divisor
// https://en.wikipedia.org/wiki/Greatest_common_divisor
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
// https://en.wikipedia.org/wiki/Greatest_common_divisor
func LCM(x, y int, numbers ...int) int {
	result := x * y / GCD(x, y)

	for i := 0; i < len(numbers); i++ {
		result = LCM(result, numbers[0])
	}

	return result
}
