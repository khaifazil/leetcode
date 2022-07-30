package main

import (
	"fmt"
)

func main() {
	str1 := "abcjust"
	str2 := "pqr"

	fmt.Println(mergeAlternately(str1, str2))
}

func mergeAlternately(word1 string, word2 string) string {
	var res string

	for i := 0; i < len(word1); i++ {
		j := i
		for j <= i {
			res += string(word1[i])
			res += string(word2[j])
			if i == len(word1)-1 && j < len(word2)-1 {
				res += word2[j+1:]
			}
			j++
		}
		if j == len(word2) && i < len(word1)-1 {
			res += word1[i+1:]
			break
		}
	}
	return res
}
