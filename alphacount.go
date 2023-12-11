package piscine

/*func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}*/

func AlphaCount(s string) int {
	srune := []rune(s)
	count := 0
	for _, value := range srune {
		if value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' {
			count++
		}
	}
	return count
}
