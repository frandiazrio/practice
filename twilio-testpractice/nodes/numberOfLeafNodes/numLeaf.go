package main

func numLeaf(root *TreeNode)int{

}


func helper(root *TreeNode, count *int)int{
	if root.left == nil && root.right{
		count++
		return count
	}

	helper(root.left, count)
	helper(root.right, count)
	return *count
}

func main(){

}