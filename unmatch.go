package piscine

func Unmatch(a []int) int {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	for 0 < len(a) {
		if len(a)-1 != 0 {
			if a[0] == a[1] {
				a = a[2:]
			} else {
				return a[0]
			}
		} else {
			return a[0]
		}
	}

	return -1
}

// For example, if a is [1, 2, 3, 4, 5], then a[2:] would result in the slice [3, 4, 5]. It includes all elements from index 2 to the end of the slice

/*package piscine

// Unmatch finds the element in the slice that does not have a corresponding pair.
// If all numbers have a corresponding pair, it returns -1.
func Unmatch(a []int) int {
	// Sorting the slice in ascending order
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	// Checking for unpaired elements
	for 0 < len(a) {
		// If there is more than one element left in the slice
		if len(a)-1 != 0 {
			// If the first two elements are equal, remove them from the slice
			if a[0] == a[1] {
				a = a[2:]
			} else {
				// If the first two elements are not equal, return the first element
				return a[0]
			}
		} else {
			// If there is only one element left in the slice, return it
			return a[0]
		}
	}

	// If no unpaired element found, return -1
	return -1
}
*/
