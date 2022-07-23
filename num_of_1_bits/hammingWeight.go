package main

import "fmt"

func main() {
	var input uint32

	input = 00000000000000000000000000001011

	fmt.Println(hammingWeight(input))
}

func hammingWeight(num uint32) int {
	sum := 0
	for num != 0 {
		sum++
		num &= (num - 1)
	}
	return sum
}
