package main

import (
	"fmt"
)

func main() {
	num := []int{1}

	fmt.Println(missingNumber(num))
}

//func missingNumber(nums []int) int {
//
//	sort.Ints(nums)
//
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i]+1 != nums[i+1] {
//			return nums[i] + 1
//		}
//	}
//
//	if nums[len(nums)-1] == len(nums) {
//		return 0
//	}
//	return len(nums)
//}

func missingNumber(nums []int) int {
	var arrsum, allsum int
	for i, v := range nums {
		arrsum += v
		allsum += i
	}
	return (allsum + len(nums)) - arrsum
}
