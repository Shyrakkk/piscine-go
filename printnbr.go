package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	SetNbr(n)
}

func SetNbr(n int) {
	z := '0'
	if n == 0 {
		z01.PrintRune(z)
		return
	}
	for i := 1; i <= n%10; i++ {
		z++
	}
	for i := -1; i <= n%10; i++ {
		z++
	}
	if n/10 != 0 {
		SetNbr(n / 10)
	}
	z01.PrintRune(z)
	return
}
