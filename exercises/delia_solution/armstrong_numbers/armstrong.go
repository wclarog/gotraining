package main

import (
	"fmt"
	"math"
)

func checkArmstrong(number int) bool {
	sum := 0
	digits := CountDigits(number)
	tmpNumber := number
	for {
		remainder := tmpNumber % 10
		sum = sum + int(math.Pow(float64(remainder), float64(digits)))
		tmpNumber /= 10
		if tmpNumber == 0 {
			break
		}
	}
	if sum == number {
		return true
	}
	return false
}

func CountDigits(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count++
	}
	return count
}

func main() {
	var number int
	fmt.Printf("Enter number to check if it's an Armstrong number:\n")
	fmt.Scanf("%d", &number)

	fmt.Println("The number", number, "is an Armstrong number: ", checkArmstrong(number))
}
