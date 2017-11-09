package main

import (
	"fmt"
	"time"
)

func getLocationTime(t time.Time, tz string) time.Time {
	loc, _ := time.LoadLocation(tz)

	locationTime, err := time.ParseInLocation("2006-01-02 15:04", t.Format("2006-01-02 15:04"), loc)
	if err != nil {
		return time.Now()
	}

	return locationTime
}

// getFlightDuration gets the duration of a flight, accounting for different airport timezones.
func getFlightDuration(depTZ, arrTZ string, dep, arr time.Time) float64 {
	oltd := getLocationTime(dep, depTZ)
	olta := getLocationTime(arr, arrTZ)

	return olta.Sub(oltd).Minutes()
}

func main() {
	format := "2006-01-02 15:04"

	dep, _ := time.Parse(format, "2017-10-11 16:50")
	arr, _ := time.Parse(format, "2017-10-11 20:50")

	depTZ := "Europe/London"
	arrTZ := "America/Jamaica"

	fmt.Println(getFlightDuration(depTZ, arrTZ, dep, arr))
	fmt.Println(arr.Sub(dep).Minutes())
}
