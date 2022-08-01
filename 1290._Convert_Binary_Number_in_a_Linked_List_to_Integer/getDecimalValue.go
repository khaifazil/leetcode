package main

import "strconv"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	currentNode := head
	temp := ""
	for currentNode != nil {
		temp += strconv.Itoa(currentNode.Val)
		currentNode = currentNode.Next
	}

	res, _ := strconv.ParseInt(temp, 2, 64)
	return int(res)
}
