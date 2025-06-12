package scale

import (
	"fmt"
	"strings"
)

var notesChromaticSharps = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
var notesChromaticFlats = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}

func chooseSharpsOrFlats(tonic string) []string {
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb":
		return notesChromaticFlats
	case "d", "g", "c", "f", "bb", "eb":
		return notesChromaticFlats
	case "C", "G", "D", "A", "E", "B", "F#":
		return notesChromaticSharps
	case "a", "e", "b", "f#", "c#", "g#", "d#":
		return notesChromaticSharps
	// need error handling here
	default:
		return notesChromaticSharps
	}
}

// Scale returns a scale based on the given tonic and interval
// String `interval` specifies the intervals uses
//
//	m for minor (semitone => shift by 1 in slice of note names)
//	M for major (whole tone => shift by 2 in slice of note names)
//	A for augmented (whole tone + semitone => shift by 3 in slice of note names)
//
// If `interval` is empty, return a chromatic scale
// If the tonic is not a valid note, return an empty slice
func Scale(tonic, interval string) []string {
	scale := []string{}
	notes := chooseSharpsOrFlats(tonic)
	upperTonic := ""
	if len(tonic) > 1 {
		tonicParts := strings.Split(tonic, "")
		upperTonic = strings.ToUpper(tonicParts[0])
		upperTonic += tonicParts[1]
	} else {
		upperTonic = strings.ToUpper(tonic)
	}
	tonicIndex := -1
	for i, note := range notes {
		if note == upperTonic {
			tonicIndex = i
			break
		}
	}
	if tonicIndex == -1 {
		fmt.Println("Invalid tonic", tonic)
		return scale
	}

	scale = append(scale, upperTonic)

	interval = strings.TrimSpace(interval)
	if interval == "" {
		for i := 1; i < 12; i++ {
			scale = append(scale, notes[(tonicIndex+i)%12])
		}
		return scale
	}
	steps := strings.Split(interval, "")
	for _, step := range steps {
		switch step {
		case "m":
			tonicIndex += 1
		case "M":
			tonicIndex += 2
		case "A":
			tonicIndex += 3
		}
		scale = append(scale, notes[tonicIndex%12])
	}
	return scale
}
