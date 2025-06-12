package main

import (
	"exercism/Go/space-age/space"
	"fmt"
)

func main() {

	seconds := 100000000
	yearsEarth := seconds / space.EarthYrSecs
	planet := space.Mars
	yearsPlanet := space.Age(float64(seconds), planet)
	fmt.Printf("Age on Earth is %d and on %s is %v", yearsEarth, planet, yearsPlanet)

}
