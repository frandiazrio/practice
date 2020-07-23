package main

import "fmt"

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func Flatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.left != nil && root.right != nil {
		return root
	}

	leftTail := Flatten(root.left)
	rightTail := Flatten(root.right)

	if leftTail != nil {
		leftTail.right = rightTail
		root.right = leftTail
		root.left = nil
	}

	if rightTail == nil {
		return leftTail
	} else {
		return rightTail
	}
}

func main() {

}
