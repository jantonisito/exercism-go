package acronym

// Source: exercism/problem-specifications
// Commit: cacf1f1 Acronym: add underscore test case (#1436)
// Problem Specifications Version: 1.7.0

type acronymTest struct {
	description string
	input       string
	expected    string
}

var testCases = []acronymTest{
	{
		description: "basic",
		input:       "Portable Network Graphics",
		expected:    "PNG",
	},
	{
		description: "lowercase words",
		input:       "Ruby on Rails",
		expected:    "ROR",
	},
	{
		description: "punctuation",
		input:       "First In, First Out",
		expected:    "FIFO",
	},
	{
		description: "all caps word",
		input:       "GNU Image Manipulation Program",
		expected:    "GIMP",
	},
	{
		description: "punctuation without whitespace",
		input:       "Complementary metal-oxide semiconductor",
		expected:    "CMOS",
	},
	{
		description: "very long abbreviation",
		input:       "Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me",
		expected:    "ROTFLSHTMDCOALM",
	},
	{
		description: "consecutive delimiters",
		input:       "Something - I made up from thin air",
		expected:    "SIMUFTA",
	},
	{
		description: "apostrophes",
		input:       "Halley's Comet",
		expected:    "HC",
	},
	{
		description: "underscore emphasis",
		input:       "The Road _Not_ Taken",
		expected:    "TRNT",
	},
}
