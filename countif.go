package piscine

/*
// tests
func main() {
	tab1 := []string{"Hello", "how", "are", "you"}
	tab2 := []string{"This", "1", "is", "4", "you"}
	answer1 := CountIf(IsNumeric, tab1)
	answer2 := CountIf(IsNumeric, tab2)
	fmt.Println(answer1)
	fmt.Println(answer2)
}
// sample function
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
*/

// func countif
func CountIf(f func(string) bool, tab []string) int {
	count := 0
	var result bool
	for _, char := range tab {
		result = f(char)
		if result == true {
			count++
		}
	}
	if count >= 1 {
		return count
	} else {
		return 0
	}
}
