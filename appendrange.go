package piscine

/*
func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}
*/
func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var ranges []int
	for i := min; i < max; i++ {
		ranges = append(ranges, i)
	}
	return ranges
}
