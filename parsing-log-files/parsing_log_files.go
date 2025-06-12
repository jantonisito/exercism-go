package parsinglogfiles

import (
	//"fmt"
	"regexp"
	"strings"
)

var prefixes []string = []string{
	"TRC",
	"DBG",
	"INF",
	"WRN",
	"ERR",
	"FTL",
}

func IsValidLine(text string) bool {
	pattern := "^\\[(" + strings.Join(prefixes, "|") + ")\\]" + ".*$"
	re := regexp.MustCompile(pattern)
	return re.Match([]byte(text))
}

func SplitLogLine(text string) []string {
	pattern := "<[~\\*=-]*>"
	re := regexp.MustCompile(pattern)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	//pattern := "(?i)\"[^\"]*password[^\"]*\"" //(?i) at the beginning to make it case insensitive
	pattern := `(?i)"[^"]*password[^"]*"` //(?i) at the beginning to make it case insensitive
	re := regexp.MustCompile(pattern)
	total := 0
	for _, line := range lines {
		total += len(re.FindAll([]byte(line), -1))
	}
	return total
}

func RemoveEndOfLineText(text string) string {
	pattern := `end-of-line\d+`
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	pattern := `User \s*([\w\d]+)`
	re := regexp.MustCompile(pattern)
	out := make([]string, len(lines))
	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			user := match[1]
			out[i] = "[USR] " + user + " " + line
		} else {
			out[i] = line
		}
	}
	return out
}
