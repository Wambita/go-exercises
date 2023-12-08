package piscine

/*import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}*/

func FirstRune(s string) rune {
	srune := []rune(s)
	return srune[0]
}
