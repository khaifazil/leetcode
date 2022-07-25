package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
	Size int
}

func (l *LinkedList) addNode(node *ListNode) {
	if l.Size == 0 {
		l.Head = node
		l.Tail = node
		l.Size++
	}

	if l.Size > 0 {
		l.Tail.Next = node
		l.Tail = l.Tail.Next
		l.Size++
	}

}

func printList(head *ListNode) {
	currentNode := head

	for currentNode != nil {
		fmt.Println(currentNode.Val)
		currentNode = currentNode.Next
	}
}

func main() {
	list := LinkedList{}

	list.addNode(&ListNode{Val: 1})
	list.addNode(&ListNode{Val: 1})
	//list.addNode(&ListNode{Val: 2})
	//list.addNode(&ListNode{Val: 3})
	//list.addNode(&ListNode{Val: 3})

	printList(list.Head)
	fmt.Println("")

	deleteDuplicates(list.Head)
	printList(list.Head)
}

//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	slow := head
//	fast := head.Next
//
//	for fast != nil {
//		if slow.Val == fast.Val {
//			slow.Next = fast.Next
//			fast = nil
//			fast = slow.Next
//		} else {
//			slow = fast
//			fast = fast.Next
//		}
//	}
//	return head
//}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else if head.Val == head.Next.Val {
		head = deleteDuplicates(head.Next)
		return head
	} else {
		head.Next = deleteDuplicates(head.Next)
		return head
	}
}
