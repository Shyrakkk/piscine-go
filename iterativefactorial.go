package piscine

func IterativeFactorial(nb int) int {

	if nb <= 1 || nb <= 20 {
		result := 1
		for a := 1; a <= nb; a++ {
			result = a * result
		}
		return result

	} else {
		return 0
	}

}
