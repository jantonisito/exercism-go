package scrabble

import "strings"

func Score(word string) int {
	if len(word) == 0 {
		return 0
	}
	score := 0
	word = strings.ToLower(strings.TrimSpace(word))
	for _, ch := range strings.Split(word, "") {
		switch ch {
		case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			score += 1
		case "d", "g":
			score += 2
		case "b", "c", "m", "p":
			score += 3
		case "f", "h", "v", "w", "y":
			score += 4
		case "k":
			score += 5
		case "j", "x":
			score += 8
		case "q", "z":
			score += 10
		}
	}
	return score
}
