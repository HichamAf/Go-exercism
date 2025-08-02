// Package weather provides tools that can forecast the current weather condition.
package weather

// CurrentCondition variable stores the current weather condition.
var CurrentCondition string

// CurrentLocation variable stores the local weather condition.
var CurrentLocation string

// Forecast function provides the weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
