package main

import "fmt"

func main() {
	arr := [][]int{{1, 2}, {3, 4}}

	fmt.Println(matrixReshape(arr, 1, 4))
}

//func matrixReshape(mat [][]int, r int, c int) [][]int {
//	if r*c != len(mat)*len(mat[0]) {
//		return mat
//	}
//
//	res := make([][]int, r)
//	for i := 0; i < r; i++ {
//		res[i] = make([]int, c)
//	}
//	count := 0
//	for i := 0; i < len(mat); i++ {
//
//		for j := 0; j < len(mat[0]); j++ {
//			res[count/c][count%c] = mat[i][j]
//			count++
//		}
//
//	}
//	return res
//}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}

	var queue []int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			queue = append(queue, mat[i][j])
		}
	}

	var res [][]int

	for i := 0; i < r; i++ {
		var temp []int
		for j := 0; j < c; j++ {
			temp = append(temp, queue[0])
			queue = queue[1:]
		}
		res = append(res, temp)
	}

	return res
}
