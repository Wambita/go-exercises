package piscine

func Unmatch(a []int) int {
	match := make(map[int]int)

	for _, num := range a {
		match[num]++
	}

	for num, count := range match {
		if count%2 != 0 {
			return num
		}
	}
	return 1
}
