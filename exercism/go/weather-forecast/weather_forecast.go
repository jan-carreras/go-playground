// Package weather gives weather forecasts.
package weather

// CurrentCondition stores the current condition.
var CurrentCondition string

// CurrentLocation stores the current location.
var CurrentLocation string

// Forecast returns the forecast for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
