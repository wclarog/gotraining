package main

import "fmt"

func isLeapYear(year int) bool {
	return (year % 4 == 0 && !(year % 100 == 0)) || year % 400 == 0
}

func setLeapYears(start, end int) []int {
	var leapYears []int

	for current := start; current <= end; current++ {
		if isLeapYear(current) {
			leapYears = append(leapYears, current)
		}
	}

	return leapYears
}

func main() {
	var start, end int

	fmt.Printf("Start year:\n")
	fmt.Scanf("%d", &start)

	fmt.Printf("End year:\n")
	fmt.Scanf("%d", &end)

	var calcLeapYears = setLeapYears(start, end)

	for _, leapYear := range calcLeapYears {
			fmt.Printf("Year %d is leap year\n", leapYear)
	}
}