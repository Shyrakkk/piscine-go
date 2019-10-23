package piscine

func LastRune(s string) rune {
	wraq := []rune(s)
	index := 0
	for index := range wraq {
		index++

	}
	return wraq[index-1]
}
