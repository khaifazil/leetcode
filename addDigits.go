package main

import "fmt"

func main() {
	num := 0
	fmt.Println(addDigits(num))
}

func addDigits(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num /= 10

		if num == 0 && res > 9 {
			num = res
			res = 0
		}
	}
	return res
}
