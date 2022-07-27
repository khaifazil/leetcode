package main

import (
	"fmt"
	"runtime"
)

func main() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}

	fmt.Println(maximumWealth(accounts))
}

//func maximumWealth(accounts [][]int) int {
//	res := 0
//
//	for _, v := range accounts {
//		temp := sum(v)
//		if res < temp {
//			res = temp
//		}
//	}
//
//	return res
//}

func sum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func maximumWealth(accounts [][]int) (max int) {
	runtime.GOMAXPROCS(len(accounts))
	calculateWealth := func(account []int, sumCh chan int) {
		sum := 0
		for _, val := range account {
			sum += val
		}
		sumCh <- sum
	}

	sumCh := make(chan int, len(accounts))
	for _, account := range accounts {
		go calculateWealth(account, sumCh)
		accountWealth := <-sumCh
		if accountWealth > max {
			max = accountWealth
		}

	}
	close(sumCh)
	return
}
