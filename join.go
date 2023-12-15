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
		words += word // adds string value to variable
		if i != len(strs)-1 {
			words += sep // adds separator to string
		}
	}
	return words
}
