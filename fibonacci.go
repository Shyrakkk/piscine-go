package piscine

func Fibonacci(index int) int {
	index := 0
	if index < 0 {
		return -1
	} else if index == 0 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}

}
