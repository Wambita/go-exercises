package piscine

/*
func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}
*/
func IsAlpha(s string) bool {
	srune := []rune(s)
	for _, value := range srune {
		if (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') || value >= '0' && value <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
