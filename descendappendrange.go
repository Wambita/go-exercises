package piscine

/*
func main() {
	fmt.Println(DescendAppendRange(10, 5))
	fmt.Println(DescendAppendRange(5, 10))
}
*/
func DescendAppendRange(max, min int) []int {
	var numb []int
	if max <= min {
		return numb
	}
	for i := max; i > min; i-- {
		numb = append(numb, i)
	}
	return numb
}
