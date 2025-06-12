package elon

import "fmt"

// TODO: define the 'Drive()' method
func (c *Car) Drive() {
	//no auto-dereferencing - no need to do (*c).battery - thanks Go!
	if c == nil || c.battery < c.batteryDrain {
		return
	}
	c.battery -= c.batteryDrain
	c.distance += c.speed
}

// TODO: define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	if c != nil {
		return fmt.Sprintf("Driven %d meters", c.distance)
	}
	return "0"
}

// TODO: define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	if c != nil {
		return fmt.Sprintf("Battery at %d%%", c.battery)
	}
	return "0"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	if c.battery <= 0 {
		return false
	}
	maxDistance := c.speed * c.battery / c.batteryDrain
	return maxDistance+c.distance >= trackDistance
}
