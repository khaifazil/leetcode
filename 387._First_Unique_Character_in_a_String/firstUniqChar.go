package main

import (
	"fmt"
)

func main() {
	s := "leetcode"

	fmt.Println(firstUniqChar(s))
}

func firstUniqChar(s string) int {
	m := make([]int, 26)
	for _, v := range s {
		m[int(v-'a')]++
	}
	for i, v := range s {
		if m[int(v-'a')] == 1 {
			return i
		}
	}
	return -1
}
