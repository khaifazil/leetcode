package main

import "fmt"

func main() {
	arr := []int{3, 2, 2, 3}
	target := 3

	result := removeElement(arr, target)
	fmt.Println(result)
}

func removeElement(nums []int, val int) int {
	ln := len(nums)
	result := ln

	slow := 0
	for fast := 0; fast < ln; fast++ {
		if nums[fast] == val {
			result--
		} else {
			nums[slow] = nums[fast]
			slow++
		}
		//fmt.Printf("loop %v, %v, result:%v\n", fast+1, nums, result)
	}
	return result
}
