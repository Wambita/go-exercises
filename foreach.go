package piscine

/*
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}

// sample func
func PrintNbr(n int) {
	fmt.Print(n)
}
*/
func ForEach(f func(int), a []int) {
	// iterate over each elem in the slice and apply the provided function
	for _, value := range a {
		f(value)
	}
}
