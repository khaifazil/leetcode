package main

import (
	"fmt"
	"net/http"
)

var nums = []int{1, 2, 3, 4, 5}
var target = 9

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int) // makes the hash map

	for currentIndex, currentNum := range nums { //range over the whole array
		//target - current number to get the difference. if current number + difference = target means pair is found
		if val, exists := hm[target-currentNum]; exists { //checks if (target-currentNum) exist in the hashmap, if true means we have found the pair
			return []int{currentIndex, val} // returns the result (val is the index of the difference number)
		}
		hm[currentNum] = currentIndex // store current number in hashmap as key with current index as value
	}

	return nil
}

func main() {

	a := twoSum(nums, target)
	fmt.Println(a)
	http.Get()

}
