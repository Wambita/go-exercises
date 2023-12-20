package main

/*
func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}
*/
func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var answer []int
	answer = make([]int, max-min)
	for i := 0; i < max-min; i++ {
		answer[i] = i + min
	}
	return answer
}
