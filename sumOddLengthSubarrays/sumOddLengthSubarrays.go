package main

import "fmt"

type Stack struct {
	s []int
}

func main() {
	arr := []int{1, 4, 2, 5, 3}
	fmt.Println(sumOddLengthSubarrays(arr))
}

func sumOddLengthSubarrays(arr []int) int {
	result := 0

	for len(arr) > 0 { //loop runs as long as there is 1 element in the stack
		result += arr[0]                      // add the first element in stack
		if len(arr)%2 == 1 && len(arr) != 1 { // if current stack len is odd & stack has more than 1 elem
			result += sumArr(arr) // add sum of stack to result
		}
		for i := 3; i < len(arr); i += 2 { // loop over odd sub arrays in stack
			result += sumArr(arr[:i]) //add sum of each subarray in current stack
		}

		arr = arr[1:] //update stack without the first elem
	}
	return result
}

func sumArr(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
