package main

import (
	gigasecond "exercism/Go/gigasecond/code"
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.String())
	//please note Second is a constant defined in `time` package
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
	// it has a value of 1e9 - so gigasecond is 1e18
	fmt.Println(float64(time.Second))

	t_plus_gsec := gigasecond.AddGigasecond(t)
	fmt.Println(t_plus_gsec.String())
	elapsed := t_plus_gsec.Sub(t)
	fmt.Println(elapsed.String())
	fmt.Println(formatDuration1(elapsed))
	fmt.Println(formatDuration2(t, elapsed))

}

// formatDuration1 returns a formatted string of the duration
// but without taking care of lap years
func formatDuration1(d time.Duration) string {
	const (
		Day  = time.Hour * 24
		Year = Day * 365
	)

	years := int(d / Year)
	d -= time.Duration(years) * Year

	days := int(d / Day)
	d -= time.Duration(days) * Day

	hours := int(d / time.Hour)
	d -= time.Duration(hours) * time.Hour

	minutes := int(d / time.Minute)
	d -= time.Duration(minutes) * time.Minute

	seconds := int(d / time.Second)

	return fmt.Sprintf("%d years, %d days, %d hours, %d minutes, %d seconds", years, days, hours, minutes, seconds)
}

// formatDuration2 returns a formatted string of the duration
// taking care of lap years
func formatDuration2(start time.Time, d time.Duration) string {
	end := start.Add(d)

	years := end.Year() - start.Year()

	start = start.AddDate(years, 0, 0)
	days := end.Sub(start).Hours() / 24

	start = start.AddDate(0, 0, int(days))
	hours := end.Sub(start).Hours()

	start = start.Add(time.Duration(hours) * time.Hour)
	minutes := end.Sub(start).Minutes()

	start = start.Add(time.Duration(minutes) * time.Minute)
	seconds := end.Sub(start).Seconds()

	return fmt.Sprintf("%d years, %.0f days, %.0f hours, %.0f minutes, %.0f seconds", years, days, hours, minutes, seconds)
}
