package main

import "fmt"

func main() {
	arr := []int{2, 2, 1}
	fmt.Println(singleNumber(arr))
}

func singleNumber(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = 0
		}
	}

	for key := range m {
		return key
	}
	return -1
}
