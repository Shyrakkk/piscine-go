package piscine

func TrimAtoi(s string) int {

	a := 0
	b := 1
	for _, i := range s {
		if i >= '0' && i <= '9' {
			c := int(i - 48)
			a = X*10 + b
		} else if i == '-' && X == 0 { //
			b = -1
		}
	}
	return a * b
}
