package piscine

/*
func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}
*/
func ConcatParams(args []string) string {
	var array string          // empty string to store arguments
	for index := range args { // loop through slice using index
		array = array + args[index] // add slice elem to arrau
		if index != len(args)-1 {   // if not last item, separate slice elem with new line
			array = array + "\n" // add new line after each elem
		}
	}
	return array
}
