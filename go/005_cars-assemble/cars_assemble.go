package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(float64(productionRate / 60) * (successRate / 100))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	rest := int(carsCount % 10)
	discount := int(carsCount - rest) / 10
	return uint( (10000 * rest) + 95000 * discount)
}

func main() {
	CalculateWorkingCarsPerHour(1547, 90)
	CalculateWorkingCarsPerMinute(1105, 90)
	CalculateCost(37)
	CalculateCost(21)
}