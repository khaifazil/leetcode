package main

import (
	"fmt"
	"math/bits"
	"sort"
)

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(sortByBits(arr))
}

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		x, y := bits.OnesCount(uint(arr[i])), bits.OnesCount(uint(arr[j]))
		if x == y {
			return arr[i] < arr[j]
		}
		return x < y
	})
	return arr
}
