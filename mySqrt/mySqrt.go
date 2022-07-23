package main

import "fmt"

func main() {
	res := mySqrt(8)
	fmt.Println(res)
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	var left int64 = 1
	var right int64 = int64(x)
	var target int64 = int64(x)
	for left < right {
		mid := left + (right-left)/2
		if mid*mid <= target && (mid+1)*(mid+1) > target {
			return int(mid)
		} else if mid*mid > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return int(left)
}
