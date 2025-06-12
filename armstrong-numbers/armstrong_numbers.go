package armstrong

import "strconv"

// intPow is a helper function to calculate the power of a number
// it is faster than using math.Pow
func IntPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	pow := len(s)
	sum := 0
	//note: s is a string, so we can iterate over it
	//no need to convert to a slice
	for _, v := range s {
		//note: v is a rune, easy way to convert it to an int
		i := int(v) - '0'
		sum += IntPow(i, pow)
	}
	return sum == n
}
