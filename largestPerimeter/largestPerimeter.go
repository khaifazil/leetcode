package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 1, 2}
	fmt.Println(largestPerimeter(arr))
}

func largestPerimeter(nums []int) int {
	sort.Ints(nums) // sort to ensure that largest is at the back

	for i := len(nums) - 1; i > 1; i-- { //loop from the back
		if (nums[i-2])+(nums[i-1]) > nums[i] { // if sides b + c > a, a non-zero triangle can be formed
			return (nums[i-2]) + (nums[i-1]) + nums[i] // return the sum of all sides
		}
	}
	return 0 // return 0 if not able to find non-zero triangle
}
