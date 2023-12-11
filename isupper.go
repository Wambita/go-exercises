package piscine

/*
func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))
}
*/
func IsUpper(s string) bool {
	srune := []rune(s)
	for _, value := range srune {
		if value >= 'A' && value <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
