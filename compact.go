package piscine

const N = 6

/*
func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}*/
//removes empty strings from a slice of strings and updates the original slice
//it takes a pointer to a slice of strings as inputs and returns the length of the updated slice

func Compact(ptr *[]string) int {
	var list []string            // empty  slice to store non empty values
	for _, value := range *ptr { // check if string value isn't empty
		if value != "" {
			list = append(list, value) // append non empty values to new slice
		}
	}
	*ptr = list      // update the original slice with new slice containing non empty values
	return len(list) // length of updated slice
}
