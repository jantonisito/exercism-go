package main

import (
	"fmt"
	"strings"
)

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// There are 6 rules to roman numerals:
// 0. Numerals are composed of chunks of symbols.
// 1. Chunks are placed from left to right in order of value, starting with the largest.
// 2. Chunks are then added together to form a number.
// 3. Rules of forming a chunk (L to R parsing):
//      a. If a smaller symbol is placed R then parsing of the chunk ends immediately w/o including it.
//      b. If a larger symbol is placed R then they form a chunk and parsing ends.
// 		c. A legal chunk can only have one smaller symbol to the left of the larger symbol.
//      d. If an equal symbols are placed to the right then it is added to the chunk
// 		e. No symbol can be repeated more than 4 times in a chunk.
// 4. Rules of forming a chunk (R to L parsing):
//		a. If a larger symbol is placed L then parsing ends of a chunk ends w/o including it.
//      b. If a smaller symbol is placed L then it is added to the chunk and parsing of a chunk ends.
// 		c. A legal chunk can only have one smaller symbol to the left of the larger symbol.
//      d. If an equal symbol is placed L then it is added to the chunk
//	  	e. No symbol can be repeated more than 4 times in a chunk

//Examples:
// 13 = XIII why not XIIV ? => if R
// 17 = XVII
// 19 = XIX  not XVIIII => [any digit cannot repeat more than 4 times]
// 39 = XXXIX => [if next digit is lower we end chunk]
// 40 = XL
// 43 = XLIIII
// 44 = XLIV
//

var r2i map[string]int = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// ParseRoman takes a roman numeral and returns a slice of strings
// representing the chunks of the roman numeral that can be be used
// to calculate the value of the roman numeral.
func ParseRoman(rNum string) ([]string, error) {
	// Validate input characters
	for _, char := range rNum {
		if _, exists := r2i[string(char)]; !exists {
			return nil, fmt.Errorf("invalid character '%c' in Roman numeral", char)
		}
	}
	rDigits := strings.Split(rNum, "")
	chunks := make([]string, 0)
	prev := 0
	chunk := ""
	endAt := len(rDigits) - 1
	for i, dig := range rDigits {
		if r2i[dig] < prev {
			if chunk != "" {
				chunks = append(chunks, chunk)
			}
			prev = r2i[dig]
			chunk = dig
			continue
		}
		if r2i[dig] == prev {
			chunk += dig
			if i == endAt {
				chunks = append(chunks, chunk)
				break
			}
			continue
		}
		if r2i[dig] > prev {
			chunk += dig
			if prev > 0 {
				chunks = append(chunks, chunk)
				chunk = ""
			} else {
				prev = r2i[dig]
			}
			continue
		}
	}
	return chunks, nil
}

// RomanToInt converts a Roman numeral string to an integer.
// It returns an error if the input string contains invalid characters.
// The function processes the Roman numeral from right to left,
// applying the rules of Roman numeral conversion.
// It handles cases where a smaller numeral precedes a larger one,
// indicating subtraction, and accumulates the total value accordingly.
func RomanToInt(rNum string) (int, error) {
	// Validate input characters
	for _, char := range rNum {
		if _, exists := r2i[string(char)]; !exists {
			return 0, fmt.Errorf("invalid character '%c' in Roman numeral", char)
		}
	}
	// check for empty string
	if rNum == "" {
		return 0, fmt.Errorf("empty string is not a valid Roman numeral")
	}
	// check for invalid length
	if len(rNum) > 15 {
		return 0, fmt.Errorf("Roman numeral is too long, maximum length is 15 characters")
	}
	// check for invalid sequences:
	// 1. no more than 3 of the same character in a row
	// 2. after a smaller character, there can be no more than one larger character
	// 3. no more than one smaller character before a larger character

	total := 0
	prevValue := 0
	for i := len(rNum) - 1; i >= 0; i-- {
		value := r2i[string(rNum[i])]
		if value < prevValue {
			total -= value
		} else {
			total += value
		}
		prevValue = value
	}
	return total, nil
}

func main() {
	// test := "XVMM"
	// test := "MCMXC"
	// test := "XCM"
	test := "MMVIIII"
	fmt.Println("Parsing roman numeral:", test)
	chunks, err := ParseRoman(test)
	if err != nil {
		fmt.Println("Error parsing roman numeral:", err)
	} else {
		fmt.Println("Chunks of the roman numeral:", chunks)
	}
	value, err := RomanToInt(test)
	if err != nil {
		fmt.Println("Error converting roman numeral to integer:", err)
	} else {
		fmt.Println("Value of the roman numeral:", value)
	}

}
