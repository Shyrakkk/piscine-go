package piscine

func swap(a *int, b *int) {

	c := *a
	*a = *b
	*b = c

}
