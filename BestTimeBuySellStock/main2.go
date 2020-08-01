package Best

import "strconv"

type BTreeNode struct {
	val         int
	left, right *BTreeNode
}

func Max(str1, str2 string) string {
	i1, _ := strconv.Atoi(str1)
	i2, _ := strconv.Atoi(str2)

	if i1 > i2 {
		return str1
	}

	return str2
}

func helper(root *BTreeNode, path string) string {
	if root.left == nil && root.right == nil {
		return path
	}

	path += strconv.Itoa(root.val)

	right := helper(root.right, path)
	left := helper(root.left, path)

	path = Max(right, left)

	return path
}

func SumRootLeaf(root *BTreeNode) int {
	if root == nil {
		return 0
	}
	path := ""
	total := helper(root, path)
	res, _ := strconv.Atoi(total)
	return res
}
