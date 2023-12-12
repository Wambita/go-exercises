package piscine

/*
func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}
*/
func IsPrintable(s string) bool {
	srune := []rune(s)
	for _, value := range srune {
		if value >= ' ' && value <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}
