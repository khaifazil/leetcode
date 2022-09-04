package main

import "fmt"

func main() {
	str := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(str)
	reverseString(str)
	fmt.Println(str)
}

func reverseString(s []byte) {
	l := len(s)
	mid := l/2 + (l % 2)

	for i := 0; i < mid; i++ {
		j := l - 1 - i
		s[i], s[j] = s[j], s[i]
	}
}
