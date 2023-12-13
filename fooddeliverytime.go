package piscine

/*
func main() {
	fmt.Println(FoodDeliveryTime("burger"))
	fmt.Println(FoodDeliveryTime("soda"))
	fmt.Println(FoodDeliveryTime("chips"))
	fmt.Println(FoodDeliveryTime("nuggets"))
	fmt.Println(FoodDeliveryTime("burger") + FoodDeliveryTime("chips") + FoodDeliveryTime("nuggets"))
}
*/
type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	burger := food{preptime: 15}
	chips := food{preptime: 10}
	nuggets := food{preptime: 12}

	if order == "burger" {
		return burger.preptime
	} else if order == "chips" {
		return chips.preptime
	} else if order == "nuggets" {
		return nuggets.preptime
	} else {
		return 404
	}
}
