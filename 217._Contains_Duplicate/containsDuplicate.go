package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 1}

	fmt.Println(containsDuplicate(arr))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})

	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}
