package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
	}

	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}

	fmt.Println(isSymmetric(root))

}

//func isSymmetric(root *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//
//	left := inorderTraversal(root.Left)
//	right := reverseInorderTraversal(root.Right)
//
//	if len(left) != len(right) {
//		return false
//	}
//
//	for i := 0; i < len(left); i++ {
//		if left[i] != right[i] {
//			return false
//		}
//	}
//	return true
//}
//
////recursive
//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	result := []int{}
//
//	left := inorderTraversal(root.Left)
//	right := inorderTraversal(root.Right)
//
//	result = append(result, left...)
//	result = append(result, root.Val)
//	result = append(result, right...)
//	return result
//}
//
//func reverseInorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	result := []int{}
//
//	left := inorderTraversal(root.Left)
//	right := inorderTraversal(root.Right)
//
//	result = append(result, right...)
//	result = append(result, root.Val)
//	result = append(result, left...)
//	return result
//}

func isSymmetric(root *TreeNode) bool {
	bfsQ := []*TreeNode{}
	bfsQ = append(bfsQ, root)
	for len(bfsQ) > 0 {
		//check palindrome
		for i := 0; i < len(bfsQ)/2; i++ {
			if bfsQ[i] == nil && bfsQ[len(bfsQ)-1-i] == nil {
				continue
			}
			if bfsQ[i] == nil || bfsQ[len(bfsQ)-1-i] == nil {
				return false
			}
			if bfsQ[i].Val != bfsQ[len(bfsQ)-1-i].Val {
				return false
			}
		}
		for _, n := range bfsQ {
			if n != nil {
				bfsQ = append(bfsQ, n.Left)
				bfsQ = append(bfsQ, n.Right)
			}
			bfsQ = bfsQ[1:]
		}
	}
	return true
}
