package piscine

/*
func main() {
	slice := []string{"Pineapple", "Honey", "Mushroom", "Tea", "Pepper", "Milk"}
	fmt.Println(ShoppingListSort(slice))
}
*/
func ShoppingListSort(slice []string) []string {
	list := make([]string, 10)
	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if len(slice[i]) > len(slice[j]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	for i, item := range slice {
		list[i] = item
	}
	return list
}
