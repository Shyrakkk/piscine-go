package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i, arg := range arg {
		if i == 0 {
			continue
		}
		for _, a := range arg {
			z01.PrintRune(a)
		}
		z01.PrintRune(10)
	}
}
