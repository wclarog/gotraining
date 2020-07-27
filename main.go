package main

import (
	"./exercises/practice2"
	"fmt"
)

func main() {
	fmt.Println("Leap years")
	practice2.PrintLeapYears(1984, 2000)
	fmt.Println("The teenager")
	practice2.TestAnwers()
	fmt.Println("Armstrong number")
	practice2.TestArmstrongNumber()
}
