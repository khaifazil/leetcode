package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "LOVELY"

	fmt.Println(toLowerCase(s))
}

func toLowerCase(s string) string {
	res := ""

	for _, v := range s {
		if !unicode.IsLower(v) && unicode.IsLetter(v) {
			v += 32
		}
		res += string(v)
	}
	return res
}
