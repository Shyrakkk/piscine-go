package piscine

func RecursivePower(nb int, power int) int {
	var result int
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	if power > 1 {
		result := nb * RecursivePower(power-1, nb-1)

	}
	return result

}
