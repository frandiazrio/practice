package main
import "fmt"

type TreeNode struct{
	val string
	left, right *TreeNode
}


func LevelOrderTraversal(root *TreeNode)[][]string{
	res := [][]string{}
	helperLevelOrderTraversal(root,0, res)
	return res
}


func helperLevelOrderTraversal(root *TreeNode, level int, res [][]string)[][]string{
	if root == nil{
		if len(res)-1 < level{
			res = append(res, []string{"null"})
		}else{
			res[level] = append(res[level], "null")
		}

	}else{
		if len(res)-1 < level{
			res = append(res, []string{root.val})
		}else{
			res[level] = append(res[level], root.val)
		}
		level++
		res = helperLevelOrderTraversal(root.left,level, res)
		res = helperLevelOrderTraversal(root.right,level, res)
	}
	return res

}

func Add(c []int, x int)[]int{
	c = append(c, x)
	return c
}

func main(){
	c := []int{1,2,3}
	c = Add(c, 4)
	fmt.Println(c)
}