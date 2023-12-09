package piscine

/*

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
*/
func LastRune(s string) rune {
	srune := []rune(s)
	return (srune[len(s)-1])
	return ' '
}
