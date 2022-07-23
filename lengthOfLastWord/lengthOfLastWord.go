package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(str))
}

func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")
	fmt.Println(arr)
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != "" {
			return len(arr[i])
		}
	}
	return 0
}

//func lengthOfLastWord(s string) int {
//	s = strings.TrimSpace(s)
//	result := strings.Split(s, " ")
//	return len(result[len(result)-1])
//}
