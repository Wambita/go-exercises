package piscine

/*
func main() {
	fmt.Println(IsLower("HELLO"))
	fmt.Println(IsLower("HELLO!"))
	fmt.Println(IsLower("fana"))
}
*/
func IsLower(s string) bool {
	srune := []rune(s)
	for _, value := range srune {
		if value >= 'a' && value <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
