package main

import "fmt"

var secret = 7

func main() {
	fmt.Println(guessNumber(40))
}

func guess(n int) int {
	if n > secret {
		return -1
	} else if n < secret {
		return 1
	}
	return 0
}

func guessNumber(n int) int {
	low := 1
	high := n
	for low <= high {
		mid := low + (high-low)/2
		n := guess(mid)
		if n == 0 {
			return mid
		} else if n == -1 {
			high = mid - 1
		} else if n == 1 {
			low = mid + 1
		}
	}
	return -1
}
