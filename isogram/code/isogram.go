// Package isogram implements check if string is an ISOGRAM
// i.e. if all non-white-space characters are unique
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram check is string is an isogram (no repeated chars)
// Simple map solution
func IsIsogram1(s string) bool {
	//hash based solution
	seen := make(map[rune]bool)
	for _, ch := range strings.ToLower(s) {
		if unicode.IsLetter(ch) {
			if seen[ch] {
				return false
			}
			seen[ch] = true
		}
	}
	return true
}

// IsIsogram check is string is an isogram (no repeated chars)
func IsIsogram2(phrase string) bool {
	phrase = strings.ToLower(phrase)

	for i, char := range phrase {
		if unicode.IsLetter(char) && strings.ContainsRune(phrase[:i], char) {
			return false
		}
	}
	return true
}

// Bobahop Bit field
// see: https://exercism.org/tracks/go/exercises/isogram/approaches/bitfield
// Simple idea where presence of Nth letter of alphabeth is recorded as flipping
// of Nth bit of 32 bit integer letterFlag (originally set to 0)
func IsIsogram3(phrase string) bool {
	theEnd := len(phrase)
	var letterFlags uint32 = 0

	var letter uint8
	var dif uint8
	for i := 0; i < theEnd; i++ {
		letter = phrase[i]

		switch {
		//a is 97, z is 122.
		case letter > 96 && letter < 123:
			dif = letter - 'a'
		//A is 65, Z is 90.
		case letter > 64 && letter < 91:
			dif = letter - 'A'
		default:
			continue
		}
		// 'a' and 'A' shift 1 0 places
		// 'z' and 'Z' shift 1 25 places
		if (letterFlags & (1 << dif)) != 0 {
			return false
		}
		// if no match we shift and flip the bit
		letterFlags |= (1 << dif)
	}
	return true
}

func IsIsogram(phrase string) bool {
	return IsIsogram1(phrase)
}
