package main

import "fmt"

func main() {
	arr1 := []int{4, 1, 2}
	arr2 := []int{1, 3, 4, 2}

	fmt.Println(nextGreaterElement(arr1, arr2))
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)

	for i, v := range nums2 {
		m[v] = i
	}

	var result []int
	for _, v := range nums1 {
		if idx, exists := m[v]; !exists || idx == len(nums2)-1 {
			result = append(result, -1)
		} else {
			//fmt.Println(idx)
			for j := idx + 1; j < len(nums2); j++ {
				//fmt.Println(num, v)
				if nums2[j] > v {
					result = append(result, nums2[j])
					break
				}
				if j == len(nums2)-1 {
					result = append(result, -1)
				}
			}

		}
	}
	return result
}
