package practice2

import "fmt"

func LeapYears( startYear, endYear int) []int {
	var years []int
	for year := startYear; year <= endYear; year++ {
		if isLeapYear(year) {
			years = append(years, year)
		}
	}
	return years
}

func isLeapYear(year int) bool {
	return (year % 4 == 0 && year % 100 != 0) || year % 400 == 0
}

func PrintLeapYears(startYear, endYear int) {
	years := LeapYears(startYear, endYear)
	for _, y := range years {
		fmt.Println(y)
	}
}