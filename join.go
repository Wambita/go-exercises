package piscine

/*
func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}
*/
func Join(strs []string, sep string) string {
	words := ""

	for i, word := range strs {
		words += word         // adds string value to variable
		if i != len(strs)-1 { // if not last word
			words += sep // adds separator to string
		}
	}
	return words
}

/*
func join2(str []string, sep string) string {
	// empty string to store final result
	result := ""

	// loop through string
	for i, word := range str {
		result += word
		if i != len(str)-1 {
			result += sep
		}
	}
	return result
}
*/
