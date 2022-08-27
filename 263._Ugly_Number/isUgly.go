package main

import "fmt"

func main() {
	n := 14

	fmt.Println(isUgly(n))
}

func isUgly(n int) bool {
	for n > 1 {
		if n%5 == 0 {
			n = n / 5
		} else if n%3 == 0 {
			n = n / 3
		} else if n%2 == 0 {
			n = n / 2
		} else {
			return false
		}
	}
	return n == 1
}
