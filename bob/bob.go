// Package bob implements Silent Bob's conversational skills
package bob

import (
	"strings"
	"unicode"
)

type replyType int

const (
	tAllElse = iota
	tQuestion
	tYell
	tYellQuestion
	tNoWords
)

type replyContent string

const (
	rQuestion     = "Sure."
	rYell         = "Whoa, chill out!"
	rYellQuestion = "Calm down, I know what I'm doing!"
	rNoWords      = "Fine. Be that way!"
	rAllElse      = "Whatever."
)

var replies = map[replyType]replyContent{
	tQuestion:     rQuestion,
	tYell:         rYell,
	tYellQuestion: rYellQuestion,
	tNoWords:      rNoWords,
	tAllElse:      rAllElse,
}

// HasLetter checks if string has any letters
func HasLetter(s string) bool {
	return strings.IndexFunc(s, unicode.IsLetter) != -1
}

// Hey picks Silent Bob's replies based on input remark
func Hey(remark string) string {
	r := strings.TrimSpace(remark)

	question := strings.HasSuffix(r, "?")
	yell := HasLetter(r) && strings.ToUpper(r) == r

	var repType replyType = tAllElse
	switch {
	case len(r) == 0:
		repType = tNoWords
	case question && yell:
		repType = tYellQuestion
	case question:
		repType = tQuestion
	case yell:
		repType = tYell
	default:
		repType = tAllElse
	}

	var reply string = string(replies[repType])
	return reply
}
