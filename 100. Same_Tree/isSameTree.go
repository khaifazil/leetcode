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
	if n == nil {
		n = &TreeNode{Val: val}
	}
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
	p := BinaryTree{Root: &TreeNode{Val: 1}}
	q := BinaryTree{Root: &TreeNode{Val: 1}}

	p.Root.Left = &TreeNode{Val: 2}
	p.Root.Right = &TreeNode{Val: 3}
	q.Root.Left = &TreeNode{Val: 2}
	q.Root.Right = &TreeNode{Val: 3}

	fmt.Println(p, q)
	fmt.Println(isSameTree(q.Root, p.Root))
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode

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

//func isSameTree(p *TreeNode, q *TreeNode) bool {
//	if (p == nil && q != nil) || (q == nil && p != nil) {
//		return false
//	}
//	res1 := inorderTraversal(p)
//	res2 := inorderTraversal(q)
//
//	if len(res1) != len(res2) {
//		return false
//	}
//
//	for i, v := range res1 {
//		if v != res2[i] {
//			return false
//		}
//	}
//	return true
//}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	qP := []*TreeNode{p}
	qQ := []*TreeNode{q}

	for len(qP) != 0 && len(qQ) != 0 {
		pNode := qP[0]
		qP = qP[1:]

		qNode := qQ[0]
		qQ = qQ[1:]

		if pNode == nil && qNode == nil {
			continue
		}
		if pNode == nil && qNode != nil || pNode != nil && qNode == nil {
			return false
		}
		if pNode.Val != qNode.Val {
			return false
		}

		qP = append(qP, pNode.Left, pNode.Right)
		qQ = append(qQ, qNode.Left, qNode.Right)
	}

	if len(qP) == 0 && len(qQ) == 0 {
		return true
	}
	return false
}
