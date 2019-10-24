package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	wyr := os.Args
	myr := 0

	for index := range wyr {
		myr = index
	}

	for i := myr; i >= 1; i-- {
		for _, bukva := range wyr[i] {
			z01.PrintRune(bukva)
		}
		z01.PrintRune('\n')
	}
}
