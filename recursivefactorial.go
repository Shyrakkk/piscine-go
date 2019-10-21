package piscine

func RecursiveFactorial(nb int) int {
	if nb >= 0 && nb <= 20 {
		return nb * RecursiveFactorial(nb-1)
	}
	if nb == 0 || nb == 1 {
		return 1

	} else {
		return 0
	}

}
