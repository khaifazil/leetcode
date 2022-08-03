package main

import "fmt"

func main() {
	arr := []int{6, 5, 4, 4}
	fmt.Println(isMonotonic(arr))
}

//func isMonotonic(nums []int) bool {
//
//	monotonicAsc := func(arr []int) bool {
//		for i := 0; i < len(arr)-1; i++ {
//			if arr[i] > arr[i+1] {
//				return false
//			}
//		}
//		return true
//	}
//
//	monotonicDsc := func(arr []int) bool {
//		for i := 0; i < len(arr)-1; i++ {
//			if arr[i] < arr[i+1] {
//				return false
//			}
//		}
//		return true
//	}
//
//	if !monotonicAsc(nums) && !monotonicDsc(nums) {
//		return false
//	}
//	return true
//}

func isMonotonic(nums []int) bool {
	isIncreasing := false
	isDecreasing := false

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] { // next number is greater than current number
			isIncreasing = true
		} else if nums[i+1] < nums[i] { // next number is less than current number
			isDecreasing = true
		}

		// monotonic array can be either decreasing OR increasing
		if isIncreasing == true && isDecreasing == true {
			return false
		}
	}

	// for arrays with all similar values, increasing and decreasing will both stay false so will return true.

	// for monotonic arrays, it will be increasing == true && decreasing == false
	// or increasing == false && decreasing == true

	return true
}
