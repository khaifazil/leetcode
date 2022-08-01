package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagara"

	fmt.Println(isAnagram(s, t))
}

//func isAnagram(s string, t string) bool {
//	m := make(map[int32]int)
//
//	for _, v := range s {
//		m[v]++
//	}
//
//	for _, v := range t {
//		if _, ok := m[v]; !ok {
//			return false
//		} else {
//			m[v]--
//			if m[v] < 0 {
//				return false
//			}
//		}
//	}
//
//	for _, v := range m {
//		if v > 0 {
//			return false
//		}
//	}
//	return true
//}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make([]int, 26)

	for i := 0; i < len(s); i++ {
		count[int(s[i])-int('a')]++
		count[int(t[i])-int('a')]--
	}

	for _, val := range count {
		if val != 0 {
			return false
		}
	}
	return true
}
