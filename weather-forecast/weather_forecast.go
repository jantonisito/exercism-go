// Package weather can forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition describes current weather condition.
var CurrentCondition string

// CurrentLocation describes current location (city) for which forecast is available.
var CurrentLocation string

// Forecast return string describing current weather forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
