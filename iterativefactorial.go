package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb < 0 {
		return 0
	}
	if nb > 0 && nb < 26 {
		result := 1
		for a := 1; a <= nb; a++ {
			result = a * result
		}
		return result

	}

}
