package piscine

func Abort(a, b, c, d, e int) int {
	num := []int{a, b, c, d, e}
	for i := 0; i <= len(num)-1; i++ {
		// sorting in ascending  order
		for j := i + 1; j < len(num); j++ {
			if num[i] > num[j] {
				num[i], num[j] = num[j], num[i]
			}
		}
	}
	// return median which is at index 2 3rd num
	return num[2]
}
