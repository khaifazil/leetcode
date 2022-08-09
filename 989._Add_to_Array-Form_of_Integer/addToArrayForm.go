package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}
	k := 516

	fmt.Println(addToArrayForm(arr, k))
	//addToArrayForm(arr, k)
}

//func addToArrayForm(num []int, k int) []int {
//	str := ""
//	for _, v := range num {
//		str += strconv.Itoa(v)
//	}
//	//fmt.Println(str)
//
//	total, _ := strconv.ParseInt(str, 10, 0)
//	fmt.Println(total)
//	//total += k
//	fmt.Println(total)
//
//	//arr := strings.Split(strconv.Itoa(total), "")
//	var res []int
//	//for _, v := range arr {
//	//	temp, _ := strconv.Atoi(v)
//	//	res = append(res, temp)
//	//}
//
//	return res
//}

func addToArrayForm(A []int, K int) []int {
	size := len(A)
	A[size-1] += K
	for i := size - 1; i > 0 && A[i] > 9; i-- {
		A[i-1] += A[i] / 10
		A[i] %= 10
	}

	if A[0] < 10 {
		return A
	}

	A0 := num2int(A[0])

	return append(A0, A[1:]...)
}

func num2int(n int) []int {
	res := make([]int, 0, 8)
	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}

	reverse(res)
	return res
}

func reverse(n []int) {
	i, j := 0, len(n)-1
	for i < j {
		n[i], n[j] = n[j], n[i]
		i++
		j--
	}
}
