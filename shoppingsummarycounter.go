package piscine

/*
func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
*/
func ShoppingSummaryCounter(str string) map[string]int {
	// split the input into a slice of words
	// list := strings.Fields(str)
	count := make(map[string]int)
	// map to store count of each
	// word to store individual words
	word := ""
	words := []string(str)
	for _, char := range words {
		if char == ' ' {
			if word != "" {
				count[word]++
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		count[word]++
	}
	// iterate over the list and store the count of each
	for _, item := range count {
		count[item]++
	}
	return count
}
