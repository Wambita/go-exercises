package piscine

/*
func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01,02,03"))
}
*/
func IsNumeric(s string) bool {
	srune := []rune(s)
	for _, value := range srune {
		if value >= '0' && value <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
