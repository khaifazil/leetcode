package main

import (
	"fmt"
)

func main() {
	str1 := "a"
	str2 := "aa"

	fmt.Println(string(findTheDifference(str1, str2)))
}

func findTheDifference(s string, t string) byte {
	Arr := make([]int, 26)

	for _, v := range s {
		charPos := v - 'a'
		Arr[charPos]++
	}
	for _, v := range t {
		charPos := v - 'a'
		if Arr[charPos]--; Arr[charPos] == -1 {
			return byte(v)
		}
	}
	return 0
}

//func findTheDifference(s string, t string) byte {
//	for _, v := range s {
//		strings.Replace(t, string(v), "", 1)
//	}
//	return t[0]
//}

//func findTheDifference(s string, t string) byte {
//	m := make(map[int32]int)
//
//	for _, v := range s {
//		m[v] = 0
//	}
//	var res int32
//	biggest := 0
//	for _, v := range t {
//		if _, ok := m[v]; !ok {
//			return byte(v)
//		}
//		m[v]++
//		if m[v] > biggest {
//			res = v
//			biggest = m[v]
//		}
//	}
//
//	return byte(res)
//}

//func findTheDifference(s string, t string) byte {
//	for _, v := range t {
//		if temp := strings.IndexRune(s, v); temp == -1 {
//			return byte(v)
//		}
//	}
//	return 0
//}
