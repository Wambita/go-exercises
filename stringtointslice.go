package piscine

/*
func main() {
	fmt.Println(StringToIntSlice("A quick brown fox jumps over the lazy dog"))
	fmt.Println(StringToIntSlice("Converted this string into an int"))
	fmt.Println(StringToIntSlice("hello THERE"))
}
*/
func StringToIntSlice(str string) []int {
	var intslice []int // empty slice var

	for _, char := range str { // loop through each slice elem
		num := int(char)                 // change char to int
		intslice = append(intslice, num) // append int to slice var
	}
	return intslice
}
