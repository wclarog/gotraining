package main

import "fmt"
import "math"

func isArmstrongNumber(val int) bool {
	textNum := fmt.Sprintf("%d", val)
	numDigits := len(textNum)
	sum := 0

	for current := val; current != 0 && sum <= val; current /= 10 {
		digit := current % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
	}

	return sum == val
}

func main() {
	/*
	for num := 1; num > 0; {
		fmt.Printf("Target:\n")
		fmt.Scanf("%d", &num)

		fmt.Printf("%d Armstrong Number condition is %v.\n", num, isArmstrongNumber(num))
	}
	*/

	for idx := 1; idx < 1000000000; idx++ {
		if isArmstrongNumber(idx) {
			fmt.Printf("%d is Armstrong Number.\n", idx)
		}
	}
}