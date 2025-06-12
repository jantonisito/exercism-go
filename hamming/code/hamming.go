// Package hamming for computing Hamming distance between strings
package hamming

import (
	"errors"
)

// Distance returns Hamming distance between strings of equal length -
// otherwise it returns -1 and raises error
func Distance(a, b string) (int, error) {
	len_a := len(a)
	if len_a != len(b) {
		return -1, errors.New("strings must have equal length to be compared")
	}
	var dist int = 0
	for i := 0; i < len_a; i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
