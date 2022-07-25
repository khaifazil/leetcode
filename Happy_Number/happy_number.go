package main

import "fmt"

func main() {
	fmt.Println(isHappy(2))
}

//floyd's cycle finding algorithm
func isHappy(n int) bool {
	x, y := n, n
	for {
		x = digitSum(x)           //tortoise
		y = digitSum(digitSum(y)) //hare
		if x == 1 || y == 1 {     // if either is 1, there is no cycle, return true
			return true
		}
		if x == y { // if either are the same, there is a cycle, return false
			return false
		}
	}
}

func digitSum(n int) int { // function to get product of all digits in n
	var digitSlice []int
	for n != 0 {
		digitSlice = append(digitSlice, n%10)
		n /= 10
	}

	var sum int
	for _, v := range digitSlice {
		sum += v * v
	}
	return sum
}
