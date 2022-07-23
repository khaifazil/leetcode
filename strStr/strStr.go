package main

import "fmt"

func main() {
	fmt.Println(strStr("heol", " "))
}

func strStr(haystack string, needle string) int {
	if needle == " " {
		return 0
	}

	ln := len(needle)
	lh := len(haystack) + 1 - ln

	for i := 0; i < lh; i++ {
		if needle == haystack[i:i+ln] {
			return i
		}
	}
	return -1
}
