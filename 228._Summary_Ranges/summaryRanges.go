package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{0, 1, 2, 4, 5, 7}

	fmt.Println(summaryRanges(arr))
}

//func summaryRanges(nums []int) []string {
//	var res []string
//	count := 0
//	for i, j := 0, 0; i < len(nums); {
//		if j >= len(nums) && count >= 1 {
//			res = append(res, strconv.Itoa(nums[i]))
//			i++
//			continue
//		}
//		if nums[i]+count != nums[j] {
//			if j-i < 2 {
//				res = append(res, strconv.Itoa(nums[i]))
//			} else {
//				temp := strconv.Itoa(nums[i]) + "->" + strconv.Itoa(nums[j-1])
//				res = append(res, temp)
//			}
//			i = j
//			count = 1
//		} else {
//			count++
//		}
//		j++
//	}
//
//	return res
//}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	var res []string
	var head int = 0

	for i := range nums {
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			continue
		}
		if head == i {
			res = append(res, strconv.Itoa(nums[i]))
		} else {
			tmp := strconv.Itoa(nums[head]) + "->" + strconv.Itoa(nums[i])
			res = append(res, tmp)
		}

		head = i + 1
	}

	return res
}
