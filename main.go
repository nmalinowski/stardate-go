package main

import (
	"fmt"
	"math"
	"time"
)

const STAR_TREK_EPOCH int = 2323 // define STAR_TREK_EPOCH as a constant

func ConvertToStardate(date time.Time) float64 {
	isLeapYear := func(year int) bool {
		return time.Date(year, time.February, 29, 0, 0, 0, 0, time.UTC).Month() == time.February
	}

	daysInYear := func(year int) int {
		if isLeapYear(year) {
			return 366
		}
		return 365
	}

	dayOfYear := func(year, month, day int) int {
		dayOfYear := []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}[month-1] + day - 1

		if month >= 2 && isLeapYear(year) {
			dayOfYear++
		}
		return dayOfYear
	}

	starYear := func(year int) float64 {
		return 1000 * float64(year-STAR_TREK_EPOCH)
	}

	starDay := func(year, month, day int) float64 {
		return 1000 / float64(daysInYear(year)) * float64(dayOfYear(year, month, day))
	}

	round := func(number float64) float64 {
		return math.Round(number*100) / 100
	}

	year, month, day := date.Date()
	return round(starYear(year) + starDay(year, int(month), day+1))
}

func main() {
	// Create a time.Time value for a specific date
	date := time.Date(2023, time.December, 14, 0, 0, 0, 0, time.UTC)

	// Call the ConvertToStardate function and print the result
	fmt.Println(ConvertToStardate(date))
}
