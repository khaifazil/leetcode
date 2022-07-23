package main

import "fmt"

func main() {
	strs := []string{"flower", "flow", "flight"}

	result := longestCommonPrefix(strs)

	fmt.Println(result)
}

//func longestCommonPrefix(strs []string) string {
//	firstString := strs[0]
//	result := ""
//	isTrue := true
//
//	index := 0
//	for isTrue {
//		var char byte
//		for _, str := range strs[1:] {
//
//			if str[index] != firstString[index] {
//				isTrue = false
//				break
//			}
//			char = str[index]
//		}
//		if isTrue {
//			result += string(char)
//		}
//		index++
//	}
//	return result
//}

func longestCommonPrefix(strs []string) string {
	var result = ""
	for i := 0; ; i++ { // init continuous loop that increments i
		for _, str := range strs { // range over strs array
			if i == len(str) || strs[0][i] != str[i] { // if length of str is 0 exit loop or we've reached the last char of 1st string, or current char is not equal to char of first string, return result and exit loop.
				return result
			}
		}
		result += string(strs[0][i]) // add current char if reached this point
	}
}
