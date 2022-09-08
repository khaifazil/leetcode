package main

import (
	"fmt"
	"strings"
)

func main() {
	pattern := "abba"
	s := "dog cat cat dog"

	fmt.Println(wordPattern(pattern, s))
}

func wordPattern(pattern string, s string) bool {
	strArr := strings.Split(s, " ")
	used := make([]bool, 26)
	if len(pattern) != len(strArr) {
		return false
	}

	m := make(map[string]byte)

	for i, v := range strArr {
		if key, ok := m[v]; !ok {
			if used[pattern[i]-'a'] {
				return false
			}
			used[pattern[i]-'a'] = true
			m[v] = pattern[i]
		} else {
			if pattern[i] != key {
				return false
			}
		}
	}
	return true
}
