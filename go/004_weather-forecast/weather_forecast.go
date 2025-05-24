// Package weather provides the weather.
package weather

// CurrentCondition returns a string.
var CurrentCondition string
// CurrentLocation return a string.
var CurrentLocation string

// Forecast is a function what gives back some information.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
