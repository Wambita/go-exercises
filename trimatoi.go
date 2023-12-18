package piscine

/*
func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
*/
func Atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	result := 0
	isNegative := false

	// check for negative sign at the start
	if s[0] == '-' {
		isNegative = true
		s = s[1:] // skip the -ve sign
	}
	// iterate through each char in the string
	for _, char := range s {
		// check if char is a digit
		if char < '0' || char > '9' {
			return 0, false
		}
		// convert digit to num value
		digit := int(char - '0')

		// update result by multiplying by 10 and adding the digit
		result = result*10 + digit
	}
	// apply the sign based on isNegative
	if isNegative {
		result = -result
	}
	return result, true
}

func TrimAtoi(s string) int {
	// trims a string and converts the remaining numeric chars to int

	// check if input string is empty, return 0
	if len(s) == 0 {
		return 0
	}

	// empty string to store numeric chars
	var num string = ""

	// loop through each char in string range for loop
	for _, n := range s {
		// negative num: if char is - and string is empty add it to string
		if n == '-' && len(num) == 0 {
			num += "-"
		}
		// skip all chars that aren't digits
		if n < '0' || n > '9' {
			continue
		}
		// add num chars to string
		num += string(n)
	}
	// check if num is empty or only has a- , return 0 if true
	if len(num) == 0 || num == "-" {
		return 0
	}
	result, ok := Atoi(num)
	if !ok {
		return 0
	}
	// convert string to int
	return result
}
