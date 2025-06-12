package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	id = strings.TrimSpace(id)

	var digits []int
	for _, char := range id {
		if unicode.IsDigit(char) {
			digits = append(digits, int(char-'0'))
		} else {
			if !unicode.IsSpace(char) {
				return false
			}
		}
	}
	numDigits := len(digits)
	if numDigits < 2 {
		return false
	}

	for i := numDigits - 1; i >= 0; i-- {
		val := digits[i]
		if (numDigits-i)%2 == 0 {
			if 2*val < 9 {
				digits[i] = 2 * val
			} else {
				digits[i] = 2*val - 9
			}
		}
	}
	var lunh int
	for _, val := range digits {
		lunh += val
	}
	return lunh%10 == 0
}
