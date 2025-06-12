package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	if m < 0 {
		h += m/60 - 1
		m = 60 + m%60
	}
	if h < 0 {
		h = 24 + h%24
	}
	hour := (h + m/60) % 24
	minute := m % 60
	return Clock{hour: hour, minute: minute}
}

func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
