package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkIsLeapYear(year int) bool {
	leap := false
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				leap = true
			} else {
				leap = false
			}
		} else {
			leap = true
		}
	}
	return leap
}

func leapNumbersGenerator(start int, end int) []int {
	if end < start {
		return []int{}
	}
	a := make([]int, 0, (end-start)+1)
	step := start
	for step <= end {
		if checkIsLeapYear(step) {
			a = append(a, step)
		}
		step += 1
	}
	return a
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter start year: ")
	start, _ := reader.ReadString('\n')
	start = strings.TrimSpace(start)

	fmt.Print("Enter end year: ")
	end, _ := reader.ReadString('\n')
	end = strings.TrimSpace(end)

	startNumber, errStart := strconv.Atoi(start)
	if errStart != nil {
		fmt.Println("invalid start year", start)
	}

	endNumber, errEnd := strconv.Atoi(end)
	if errEnd != nil {
		fmt.Println("invalid end year", endNumber)
	}

	if errStart == nil && errEnd == nil {
		fmt.Println("Leap numbers between ", start, "and ", end, "years are: ")
		fmt.Println(leapNumbersGenerator(startNumber, endNumber))
	}
}
