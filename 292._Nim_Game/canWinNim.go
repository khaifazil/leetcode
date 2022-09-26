package main

import "fmt"

func main() {
	n := 1

	fmt.Println(canWinNim(n))
}

func canWinNim(n int) bool {
	return n%4 != 0
}
