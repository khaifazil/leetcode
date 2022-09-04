package main

import (
	"fmt"
)

func main() {
	s := "aA"

	fmt.Println(reverseVowels(s))
}

func reverseVowels(s string) string {
	res := []byte(s)

	left := 0
	right := len(res) - 1

	for right > left {

		if isVowel(res[left]) && isVowel(res[right]) {
			res[left], res[right] = res[right], res[left]
			left++
			right--
			continue
		}
		if isVowel(res[left]) && !isVowel(res[right]) {
			right--
			continue
		}
		left++
	}

	return string(res)
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'o' || b == 'i' || b == 'u' || b == 'A' || b == 'E' || b == 'O' || b == 'I' || b == 'U'
}
