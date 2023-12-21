package piscine

/*
// test
func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(IsNumeric, a1)
	result2 := Any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
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
// func any
func Any(f func(string) bool, a []string) bool {
	count := 0
	var result bool
	for _, char := range a {
		result = f(char)
		if result == true {
			count++
		}
	}
	if count >= 1 {
		return true
	} else {
		return false
	}
}
