package main

import (
	"fmt"
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

func main() {
	//var pattern string = `^\[(TRC|DBG|INF|WRN|ERR|FTL)\].*$`
	pattern := "^\\[(" + strings.Join(prefixes, "|") + ")\\]" + ".*$"
	fmt.Println(pattern)
	re := regexp.MustCompile(pattern)
	//re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\].+$`)
	test1 := "Any old [ERR] text"
	test2 := "[INF]Any old [ERR] text"

	fmt.Println(re.Match([]byte(test1)))
	fmt.Println(regexp.MatchString(pattern, test1))
	fmt.Println(re.Match([]byte(test2)))
	fmt.Println(regexp.MatchString(pattern, test2))

	pattern = "\"[^\"]*password[^\"]*\""
	re = regexp.MustCompile(pattern)
	tests := []string{
		`[INF] passWord`,
		`"passWord"`,
		`[INF] User saw error message "Unexpected Error" on page load.`,
		`[INF] The message "Please reset your password" was ignored by the user`,
	}
	for _, t := range tests {
		res := re.FindAll([]byte(t), -1)
		fmt.Println(len(res), t)
	}

}
