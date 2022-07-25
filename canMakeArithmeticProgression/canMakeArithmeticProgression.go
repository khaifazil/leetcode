package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 5, 1}
	fmt.Println(canMakeArithmeticProgression(arr))
}

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) < 2 {
		return true
	}

	sort.Ints(arr)

	diff := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != diff {
			return false
		}
	}
	return true
}
