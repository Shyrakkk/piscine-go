package piscine

func LastRune(s string) rune {
	wraq := []rune(s)
	for index := range wraq {
		index++

		return wraq[index-1]

	}
}
