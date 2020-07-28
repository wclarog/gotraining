package main

import "fmt"

func IsLeapYear(year int) bool{
	if year % 400 == 0 && year % 100 == 0{
		return true
	}else if year % 100 == 0{
		return false
	}else if year % 4 == 0{
		return true
	} else {
		return false
	}
}

func GetLeapYears(start int, end int) []int{
	var leapYears []int

	if end >= start{
		currentYear := start
		for currentYear <= end{
			if IsLeapYear(currentYear){
				leapYears = append(leapYears, currentYear)
			}
			currentYear++
		}
	}

	return leapYears
}

func main() {
	a := GetLeapYears(1980, 2000)
	fmt.Println("Print leap years between 1980 and 2000")
	fmt.Println(a)
	fmt.Println("The result should be 1980, 1984, 1988, 1992, 1996, 2000")
}