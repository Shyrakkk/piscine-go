package piscine

func AlphaCount(str string) int {
	count := 0

	for _, i := range str {
		if i >= 'a' && i <= 'z' || i >= (65) && i <= (90) {
			count++
		}
	}
	return count

}
