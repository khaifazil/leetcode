package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{3, 1, 1, 2}

	fmt.Println(thirdMax(input))
}

func thirdMax(nums []int) int {

	max, second, third := math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {
		if v == max || v == second || v == third {
			continue
		}

		switch {
		case v > max:
			max, second, third = v, max, second
		case v > second:
			second, third = v, second
		case v > third:
			third = v
		}
	}

	if third == math.MinInt64 {
		return max
	}

	return third
}
