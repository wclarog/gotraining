package main

import (
	"fmt"
	"math"
)

func checkArmstrong(number int) bool {
	sum := 0
	digits := CountDigits(number)
	for tmpNumber := number; tmpNumber != 0; tmpNumber /= 10 {
		remainder := tmpNumber % 10
		sum = sum + int(math.Pow(float64(remainder), float64(digits)))
	}
	return sum == number
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

	fmt.Println("The number ", number, "is an Armstrong number: ", checkArmstrong(number))
}
