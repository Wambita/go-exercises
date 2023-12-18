package piscine

import (
	"strings"
)

/*
func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
*/
func ShoppingSummaryCounter(str string) map[string]int {
	// empty slice to store count of each string
	count := make(map[string]int)

	// split text into slices of words
	words := strings.Fields(str)

	// iterate through the string and store count of each item
	for _, item := range words {
		count[item]++
	}
	return count
}
