package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func (n *TreeNode) insert(val int) {
	if val < n.Val {
		if n.Left == nil {
			n.Left = &TreeNode{Val: val}
		} else {
			n.Left.insert(val)
		}
	} else if val > n.Val {
		if n.Right == nil {
			n.Right = &TreeNode{Val: val}
		} else {
			n.Right.insert(val)
		}
	}
}

func main() {
	root := &TreeNode{
		Val: 1,
	}

	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Println(inorderTraversal(root))
}

//recursive
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

//iterative approach
//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	var result []int
//	currentNode := root
//	for currentNode != nil {
//		if currentNode.Left == nil {
//			result = append(result, currentNode.Val)
//			currentNode = currentNode.Right
//			continue
//		}
//
//		child := currentNode.Left
//		for child.Right != nil {
//			child = child.Right
//		}
//
//		child.Right = currentNode
//		currentNode, currentNode.Left = currentNode.Left, nil
//	}
//	return result
//}

//stack approach
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)
		root = root.Right
	}

	return ans
}
