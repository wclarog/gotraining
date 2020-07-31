package practice2

import (
	"fmt"
	"math"
)

func IsAnArmstrongNumber(x int) bool {
	var armstrongNumber = 0.0
	numStr := string(x)
	power := len(numStr)
	for _, d := range numStr {
		armstrongNumber += math.Pow(float64(d), float64(power))
	}
	return float64(x) == armstrongNumber
}

func TestArmstrongNumber() {
	nums := [4]int{9, 10, 153, 154}
	for _, n := range nums {
		if IsAnArmstrongNumber(n) {
			fmt.Printf("%d is an Armstrong number\n", n)
		} else {
			fmt.Printf("%d is not an Armstrong number\n", n)
		}
	}
}