package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{1, 2},
		{3, 1},
		{2, 4},
		{2, 3},
		{4, 4},
	}

	fmt.Println(arr)

	x := 3
	y := 4

	fmt.Println(nearestValidPoint(x, y, arr))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func nearestValidPoint(x int, y int, points [][]int) int {
	result := -1
	dist := math.MaxInt32

	//loop over 2d array
	for i, v := range points {
		//search for same x or y positions
		if v[0] == x || v[1] == y {
			//if found apply manhattan formula
			current := abs(v[0]-x) + abs(v[1]-y)
			fmt.Printf("loop %v, current: %v\n", i, current)
			//check if smaller than current result
			if current < dist {
				//save smaller
				result = i
				dist = current
			}
		}
	}
	//return result at the end
	return result
}
