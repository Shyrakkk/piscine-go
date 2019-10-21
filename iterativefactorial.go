package piscine

func IterativeFactorial(nb int) int {

	if nb <= 1 || nb > 20 {
		return 0
	} else {
		result := 1
		a := 1
		for ; a <= nb; a++ {
			result = result * a
		}
		return result
	}

}
