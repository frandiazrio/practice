package main

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func insert(root *TreeNode, k int) *TreeNode {
	if root == nil {
		return nil
	}
	if k == root.val {
		return root
	} else if k < root.val {
		if root.left != nil {
			insert(root.left, k)
		} else {
			root.left = &TreeNode{
				val:   k,
				left:  nil,
				right: nil,
			}
		}
	} else if k > root.val {
		if root.right != nil {
			insert(root.right, k)
		}
	} else {

	}
}
