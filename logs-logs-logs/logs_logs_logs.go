package logs

import (
	"unicode/utf8"
)

var recommendation = '\u2757'

// Note Unicode characters aoutside of Basic Multilingual Plane (BMP)
// cannot be specifies as rune literals like '\u1F5D0' but can be converted
// to a rune via: uft8.DecodeRuneDecodeRuneInString.
// For same BMP reason "\u1F5D0" does not work - only U+0000 to U+FFFF
// code strings can be represented and decoded. "\u1F5D0" is out of the range.
// but sumple trick like fmt.Printf("% x\n", search) converts it to a string
// of hexadecimal numbers:
var search = "\xf0\x9f\x94\x8d"
var weather = '\u2600'

// Application identifies the application emitting the given log.
func Application(log string) string {
	search_rune, _ := utf8.DecodeRuneInString(search)
	for _, char := range log {
		switch char {
		case recommendation:
			return "recommendation"
		case search_rune:
			return "search"
		case weather:
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := []rune{}
	for _, char := range log {
		if char == oldRune {
			newLog = append(newLog, newRune)
		} else {
			newLog = append(newLog, char)
		}
	}
	return string(newLog)
}

// note that the version below will not work
// when newLog is created via `make` it is filled with \x00 runes
// but when you iterate via `range log` the loop moves correctly
// from rune to rune and if rune is wider than the index will skip
// thus leaving `\x00` in newLog
// func Replace(log string, oldRune, newRune rune) string {
// 	newLog := make([]rune, len(log))
// 	for i, char := range log {
// 		if char == oldRune {
// 			newLog[i] = newRune
// 		} else {
// 			newLog[i] = char
// 		}
// 	}
// 	return string(newLog)
// }

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	count := utf8.RuneCountInString(log)
	return count <= limit
}
