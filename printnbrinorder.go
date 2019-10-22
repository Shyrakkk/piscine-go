package piscine

import "github.com/01-edu/z01"

func main() {
	PrintNbrInOrder(32132)
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n > 0 {
		var arsh []int
		eachValue := 0

		arshCount := 0
		var minValue int
		for n != 0 {
			eachValue = n % 10
			n /= 10
			arsh = append(arsh, eachValue)
		}

		for count := range arsh {
			arshCount = count + 1
		}
		for i := 0; i < arshCount-1; i++ {
			for j := 0; j < arshCount-i-1; j++ {
				if arsh[j] > arsh[j+1] {
					minValue = arsh[j]
					arsh[j] = arsh[j+1]
					arsh[j+1] = minValue
				}
			}
		}
		for i := 0; i < arshCount; i++ {
			z01.PrintRune(rune(arsh[i] + 48))
		}
	}
}
