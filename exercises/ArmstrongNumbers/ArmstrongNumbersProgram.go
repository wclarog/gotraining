package main

import (
	"fmt"
	"math"
	"strconv"
)

func IsArmstrongNumber(number int) bool{
	numberStr := fmt.Sprintf("%d", number)
	pow := len(numberStr)

	if pow == 1{
		return 0 != number
	} else{
		sum := 0
		for i := range numberStr{
			current := string(numberStr[i])
			j, _ := strconv.ParseInt(current, 10, 64)
			sum += int(math.Pow(float64(j), float64(pow)))
		}

		return sum == number
	}
}

func main() {
	zero := IsArmstrongNumber(0)
	println(zero)
	nine := IsArmstrongNumber(9)
	println(nine)
	ten := IsArmstrongNumber(10)
	println(ten)
	hundredFiftyThree := IsArmstrongNumber(153)
	println(hundredFiftyThree)
	hundredFiftyFour := IsArmstrongNumber(154)
	println(hundredFiftyFour)
}
