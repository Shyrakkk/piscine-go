package piscine

func Fibonacci(index int) int {
	a := index - 1
	if index < 0 {
		return -1
	} else {
		return a + index

	}

}
