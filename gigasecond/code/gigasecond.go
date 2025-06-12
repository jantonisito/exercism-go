// Packgate to add gigasecond to time
package gigasecond

// import path for the time package from the standard library https://pkg.go.dev/time
import (
	"time"
)

const Gigasecond = time.Second * 1_000_000_000

//const Gigasecond = 1e18

// AddGigasecond returns time after 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	//return t.Add(time.Second * 1000000000)
	//giga := time.Duration(math.Pow(10.0, 9.0))
	//return t.Add(time.Second * giga)

	return t.Add(Gigasecond)
}
