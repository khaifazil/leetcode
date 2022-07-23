package main

import "fmt"

func main() {
	n := 4421

	fmt.Println(subtractProductAndSum(n))
}

func subtractProductAndSum(n int) int {
	product := 1
	sum := 0

	for n != 0 {
		product *= n % 10
		sum += n % 10
		n /= 10
	}

	return product - sum
}
