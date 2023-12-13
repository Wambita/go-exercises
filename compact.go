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

func Compact(ptr *[]string) int {
	var list []string
	for _, value := range *ptr {
		if value != "" {
			list = append(list, value)
		}
	}
	*ptr = list
	return len(list)
}
