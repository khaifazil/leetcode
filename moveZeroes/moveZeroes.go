package main

import "fmt"

func main() {
	nums := []int{1, 0, 1}

	fmt.Println(moveZeroes(nums))
}

func moveZeroes(nums []int) []int {
	ptr1 := 0

	for ptr2 := 0; ptr2 < len(nums); ptr2++ {
		if nums[ptr2] != 0 {
			nums[ptr1], nums[ptr2] = nums[ptr2], nums[ptr1]
			ptr1++
		}
	}
	return nums
}
