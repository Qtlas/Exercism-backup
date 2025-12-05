//Package weather implements a client for weather.com API.
package weather

var (
    //CurrentCondition represent the current condition.
	CurrentCondition string
    //CurrentLocation represent the current city.
	CurrentLocation  string
)

//Forecast the prevision for a certain city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
