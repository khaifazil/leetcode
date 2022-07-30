package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "1326#"
	fmt.Println(freqAlphabets(s))
}

func freqAlphabets(s string) string {
	var res string

	for i := len(s) - 1; i >= 0; i-- {
		var temp int
		if s[i] != '#' {
			temp, _ = strconv.Atoi(string(s[i]))
		} else {
			temp, _ = strconv.Atoi(s[i-2 : i])
			i -= 2
		}
		res = string(rune(temp+'`')) + res
	}
	return res
}
