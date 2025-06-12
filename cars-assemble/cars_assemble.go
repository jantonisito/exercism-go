package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60.0)
}

// CalculateCost works out the cost of producing the given number of cars.
// Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not.
// But with a bit of planning, 10 cars can be produced together for $95,000.
func CalculateCost(carsCount int) uint {
	return uint(95000*(carsCount/10) + 10_000*(carsCount%10))
}
