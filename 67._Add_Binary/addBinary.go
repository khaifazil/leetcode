package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := "11"
	b := "1"

	fmt.Println(addBinary(a, b))
}

//func addBinary(a string, b string) string {
//	aInt, _ := strconv.Atoi(a)
//	bInt, _ := strconv.Atoi(b)
//	sum := aInt + bInt
//
//	temp := strconv.FormatInt(int64(sum), 2)
//
//	return temp
//}

func addBinary(a string, b string) string {
	aInt, bInt, sum := new(big.Int), new(big.Int), new(big.Int)
	aInt.SetString(a, 2)
	bInt.SetString(b, 2)
	sum.Add(aInt, bInt)
	return sum.Text(2)
}
