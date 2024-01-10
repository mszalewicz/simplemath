package simplemath

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
		result = LCM(result, numbers[0])
	}

	return result
}
