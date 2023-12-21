package piscine

/*
// test fun
func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrimed, a)
	fmt.Println(result)
}

// sample func
func IsPrimed(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
		}
		return false
	}
	return true
}
*/
func Map(f func(int) bool, a []int) []bool {
	// create a slice to store the result values
	var result []bool
	// iterate over each elem and apply func
	for _, val := range a {
		result = append(result, f(val))
	}
	return result
}
