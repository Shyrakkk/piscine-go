package main
import ( 
	"github.com/01-edu/z01"
	"os"
)	
func main() {
	arg := os.Args
	for ind, i := range arg {
		if ind == 0 {
			continue 
		}
		for_, a := range i {
			z01.PrintRune(a)
		}
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
