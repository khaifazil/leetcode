package main

import "fmt"

func main() {
	s1 := "bank"
	s2 := "kanb"

	fmt.Println(areAlmostEqual(s1, s2))
}

func areAlmostEqual(s1 string, s2 string) bool {
	count := 0
	var idx1 int
	var idx2 int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
			if count > 2 {
				return false
			}

			if count-1 == 0 {
				idx1 = i
			} else {
				idx2 = i
			}
		}
	}
	if count == 0 {
		return true
	}

	if s1[idx1] == s2[idx2] && s1[idx2] == s2[idx1] {
		return true
	}
	return false
}
