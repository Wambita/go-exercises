package piscine

/*
func main() {
	fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	fmt.Println(StringToIntSlice("Converted this string into an int"))
	fmt.Println(StringToIntSlice("hello THERE"))
}
*/
func StringToIntSlice(str string) []int {
	var intslice []int

	for _, char := range str {
		num := int(char)
		intslice = append(intslice, num)
	}
	return intslice
}
