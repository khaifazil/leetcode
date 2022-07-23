package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result := removeDuplicates(arr)
	fmt.Println(result)
}

// 2 ptr solution
func removeDuplicates(nums []int) int {
	ln := len(nums)
	if ln <= 1 {
		return ln
	}
	slow := 0                          // points to  the index of last filled position
	for fast := 1; fast < ln; fast++ { // increment fast at end of each loop
		if nums[slow] != nums[fast] { // if its not a duplicate element
			slow++                  // move slow ptr to next element
			nums[slow] = nums[fast] // copy fast(current not duplicate) into slow
		}
		fmt.Printf("loop %v, %v\n", fast, nums)
	}

	return slow + 1
}

//func removeDuplicates(nums []int) int {
//	if len(nums) == 0 {
//		return 0
//	}
//	count := 1
//	currentNum := nums[0]
//	fmt.Printf("loop %v, %v\n", 0, nums)
//	for i := 1; i < len(nums); i++ {
//		if nums[i] != currentNum {
//			currentNum = nums[i]
//			count++
//		} else {
//			nums = append(nums[:i], nums[i+1:]...)
//			i--
//		}
//		fmt.Printf("loop %v, %v\n", i, nums)
//	}
//	return count
//}
