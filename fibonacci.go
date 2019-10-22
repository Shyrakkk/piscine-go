package piscine

func Fibonacci(index int) int {
	a := index - 1
	b := index - 2
	if index < 0 {
		return -1
	} else {
		return Fibonacci(a) + Fibonacci(b)

	}

}
