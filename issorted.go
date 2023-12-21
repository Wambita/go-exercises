package piscine

/*
func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{801206, 909401, -633231, -821370, 828209, -671399, 125318, -634097}
	a4 := []int{-888201, 836859, 852494, 514358, 296482, 309794, -47446, 590523}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)
	result3 := IsSorted(f, a3)
	result4 := IsSorted(f, a4)

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}

// sample func
func f(a, b int) int {
	return a - b
}
*/
// func issorted
func IsSorted(f func(a, b int) int, a []int) bool {
	var sorted bool = true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			// if slice +ve , slice not sorted return false
			sorted = false
		}
	}
	if !sorted {
		sorted = true
		for i := 0; i < len(a)-1; i++ {
			if f(a[i], a[i+1]) < 0 {
				sorted = false
			}
		}
	}

	return sorted
}
