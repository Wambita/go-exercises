package piscine

/*
func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}
*/

func BasicJoin(elems []string) string {
	var list string              // empty string to store strings
	for _, char := range elems { // loop through string to add individual elems to string
		list += char
	}
	return list
}
