package main

import "fmt"

func main() {
	arr := [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	fmt.Println(diagonalSum(arr))
}

func diagonalSum(mat [][]int) int {
	res := 0
	l := len(mat)
	for i := 0; i < l; i++ {
		res += mat[i][i] + mat[i][l-1-i]
	}
	if l%2 == 1 {
		res -= mat[l/2][l/2]
	}
	return res
}

//func diagonalSum(mat [][]int) int {
//	res := 0
//	for i := 0; i < len(mat); i++ {
//		res += mat[i][i]
//		if i != len(mat)-i-1 {
//			res += mat[i][len(mat)-i-1]
//		}
//	}
//	return res
//}
