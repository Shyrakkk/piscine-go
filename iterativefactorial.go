package piscine

func IterativeFactorial(nb int) int {

	if nb <= 1 && nb <= 20 {
		result := 1
		a := 1
		for ; a <= nb; a++ {
			result = result * a
		}
		return result
	} else {
		return 0
	}

}
