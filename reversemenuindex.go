package piscine

/*
func main() {
	fmt.Println(ReverseMenuIndex([]string{"desserts", "mains", "drinks", "starters"}))
}
*/
func ReverseMenuIndex(menu []string) []string {
	reverse := make([]string, len(menu)) //

	for j := range menu {
		reverse[len(menu)-j-1] = menu[j]
	}
	return reverse
}
