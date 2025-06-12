// Package triangle implement a funtion KindFromSides
// that classifies triangles based on lengths of sides
package triangle

import (
	"math"
	"sort"
)

// Kind is triange type
type Kind int

// Kind type constants
const (
	NaT Kind = iota
	Sca
	Iso
	Equ
)

// ValidateTriange checks if sides describe proper triangles
func ValidateTriange(a, b, c float64) bool {
	// checking NaT cases
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	if (a == 0) || (b == 0) || (c == 0) {
		return false
	}
	if (a+b < c) || (b+c < a) || (a+c < b) {
		return false
	}
	return true
}

// KindFromSides1 returns the Kind of triangle based on sides.
// Benchmark:
// goos: windows
// goarch: amd64
// pkg: exercism/Go/triangle/code
// cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
// BenchmarkKind-8           321356              3249 ns/op
func KindFromSides1(a, b, c float64) Kind {
	if !ValidateTriange(a, b, c) {
		return NaT
	}

	sides := []float64{a, b, c}
	sort.Float64s(sides)

	// if all sides equal
	if sides[0] == sides[1] && sides[1] == sides[2] {
		return Equ
	}
	// if two sides equal
	if sides[0] == sides[1] || sides[1] == sides[2] {
		return Iso
	}
	return Sca
}

// KindFromSides2 returns the Kind of triangle based on sides.
// This version eliminates call to sort.
// Benchmark:
// goos: windows
// goarch: amd64
// pkg: exercism/Go/triangle/code
// cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
// BenchmarkKind-8          1819300               632.3 ns/op
func KindFromSides2(a, b, c float64) Kind {
	if !ValidateTriange(a, b, c) {
		return NaT
	}

	count := 3
	if b == a {
		count--
		if c == a {
			count--
		}
	} else {
		if (c == a) || (c == b) {
			count--
		}
	}

	switch count {
	case 3:
		return Sca
	case 2:
		return Iso
	case 1:
		return Equ
	default:
		return NaT
	}
}

// KindFromSides3 returns the Kind of triangle based on sides.
// This version uses set.
// goos: windows
// goarch: amd64
// pkg: exercism/Go/triangle/code
// cpu: Intel(R) Core(TM) i7-10510U CPU @ 1.80GHz
// BenchmarkKind-8           358212              2832 ns/op
func KindFromSides3(a, b, c float64) Kind {
	if !ValidateTriange(a, b, c) {
		return NaT
	}

	type blank struct{}
	// code snipped
	set := make(map[float64]blank)
	set[a] = blank{}
	set[b] = blank{}
	set[c] = blank{}

	switch len(set) {
	case 3:
		return Sca
	case 2:
		return Iso
	case 1:
		return Equ
	default:
		return NaT
	}
}

// KindFromSides points to version of function currently being used
var KindFromSides = KindFromSides3
