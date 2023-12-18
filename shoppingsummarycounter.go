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
	ToSlice := func(str string) []string {
		// empty slice to store individual words
		var words []string

		// var to keep track of the start of current word and if we are inside a word
		var wordStart int
		inWord := false

		// loop through each char in the input string
		for i := 0; i < len(str); i++ {
			// chech if current is whitespace char
			isSpace := str[i] == ' '

			// if we are inside a word, check for the end of the word
			if inWord {
				if isSpace { // add to slice and reset inWord flag
					words = append(words, str[wordStart:i])
					inWord = false

				}
			} else {
				// if not inside a word check for start of word
				if !isSpace {
					// set wordStart to current index and set inword  flag to true
					wordStart = i
					inWord = true
				}
			}

		}
		// if still in a word when reaching the end of the string , add word
		if inWord {
			words = append(words, str[wordStart:])
		}
		return words
	}

	// end of function

	count := make(map[string]int)
	words := ToSlice(str)
	// iterate through the string and store count of each item
	for _, item := range words {
		count[item]++
	}
	// empty slice to store count of each string

	// split text into slices of words
	// words := strings.Fields(str)

	return count
}
