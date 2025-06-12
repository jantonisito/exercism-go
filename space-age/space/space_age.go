// Package space calculates age on diff planets (in local years)
package space

//hneeded to round up age
import "math"

// Planet is planets name
type Planet string

// Data for planets
const (
	Mercury Planet = "Mercury"
	Venus   Planet = "Venus"
	Earth   Planet = "Earth"
	Mars    Planet = "Mars"
	Jupiter Planet = "Jupiter"
	Saturn  Planet = "Saturn"
	Uranus  Planet = "Uranus"
	Neptune Planet = "Neptune"
)

// EarthYrSecs in number of secs in Earth's year
// Note the Gregorian year has 365.2425 days
const EarthYrSecs = 365.2425 * 24 * 60 * 60

// ration of length of planet year to Earth year
var yearRatio = map[Planet]float64{
	Mercury: 0.2408467,
	Venus:   0.61519726,
	Earth:   1.0,
	Mars:    1.8808158,
	Jupiter: 11.862615,
	Saturn:  29.447498,
	Uranus:  84.016846,
	Neptune: 164.79132,
}

// Age returns age on planet in local years given age in secs and the Planet type
func Age(ageInSeconds float64, planet Planet) float64 {
	var age float64
	//note if sets `ratio` and `exists (`exist` is used to check condition)
	if ratio, exists := yearRatio[planet]; exists {
		age = (ageInSeconds / EarthYrSecs) / ratio

	} else {
		return -1.0
	}
	//rounding up to 2 digits (strictly speaking not necessary
	return math.Round(age*100) / 100.0
}
