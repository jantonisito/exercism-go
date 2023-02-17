// Package acronym allows creation of acronyms from strings
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns acronym of a string split on whitespace and hyphens
func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	//see: https://pkg.go.dev/strings#Fields
	sArray := strings.Fields(s)
	acr := []rune{}
	for _, val := range sArray {
		val = strings.TrimSpace(val)
		for _, ch := range val {
			if unicode.IsLetter(ch) {
				acr = append(acr, unicode.ToUpper(ch))
				break //break after first letter
			}
		}
	}
	return string(acr)
}
