package piscine

func LastRune(s string) rune {
	for index := range wraq {
		index++
		wraq := []rune(s)
		return wraq[index-1]

	}
}
