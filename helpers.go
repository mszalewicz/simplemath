package main

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
		x = GCD(x, numbers[0])
	}

	return x
}
