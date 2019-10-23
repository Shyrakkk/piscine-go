package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)
	wrq1 := 0
	wrq2 := 0

	for range a {
		wrq1++
	}
	for range b {
		wrq2++
	}
	if wrq2 == 0 {
		return 0
	}

	for index, value := range a {
		if value == b[0] && wrq1 >= wrq2+index-1 {
			wwwr := 1
			for i := 1; i < wrq2; i++ {
				if b[i] == a[index+i] {
					wwwr++
				}
			}
			if wwwr == wrq2 {
				return index
			}
		}
	}
	return -1
}
