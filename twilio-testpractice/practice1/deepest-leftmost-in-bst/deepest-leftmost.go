package main

import "math"

func deepestLeftMost(root *TreeNode)*TreeNode{

}


func Helper(root *TreeNode, currentLevel int, isLeft bool)*TreeNode{
	if root == nil{
		return -1
	}

	if isLeft{
		//check child nodes
		if root.left == nil && root.right == nil{
			return currentLevel
		}
	}

	left := Helper(root.left, currentLevel+1, true)
	right := Helper(root.right, currentLevel+1, false)

	return math.Max(float64(left), float64(right))

}

func main(){
/*

				A
			  /   \
			 C     B
            / \   / \
               D E
*/
}