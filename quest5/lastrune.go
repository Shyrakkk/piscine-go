package piscine

func LastRune(s string) rune {
	wraq := []rune(s)
	index := 0
	for s := range wraq {
		index = s + 1

	}
	return wraq[index-1]
}
