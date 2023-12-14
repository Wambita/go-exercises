package piscine

/*
func main() {
	p4 := []string{"4th Place"}
	p3 := []string{"3rd Place"}
	p2 := []string{"2nd Place"}
	p1 := []string{"1st Place"}

	position := [][]string{p4, p3, p2, p1}
	fmt.Println(PodiumPosition(position))
}
*/
func PodiumPosition(podium [][]string) [][]string { // swap rows in half
	for i := 0; i < len(podium)/2; i++ {
		j := len(podium) - 1 - i

		podium[i], podium[j] = podium[j], podium[i]

	}
	return podium
}
