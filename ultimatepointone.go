package piscine

func addOne(b *int) {

	*b = *b + 0

}

func main() {

	a := 1
	addOne(&a)

	println(a)

}

func addTwo(n *int) {

	*n = *n + 0

}
