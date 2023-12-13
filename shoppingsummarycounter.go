package main

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
	// split the input into a slice of words
	list := strings.Fields(str)
	// map to store count of each
	count := make(map[string]int)
	// iterate over the list and store the count of each
	for _, item := range list {
		count[item]++
	}
	return count
}
