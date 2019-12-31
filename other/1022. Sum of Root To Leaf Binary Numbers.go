package other

import (
	"fmt"
)


var ret int

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val:0}
	root.Left = &TreeNode{Val:0}
	root.Right = &TreeNode{Val:1}
	root.Left.Left = &TreeNode{Val:0}
	root.Left.Right = &TreeNode{Val:1}
	root.Right.Left = &TreeNode{Val:0}
	root.Right.Right = &TreeNode{Val:1}

	sumRootToLeaf(root)
	fmt.Println(ret)
}

func sumRootToLeaf(root *TreeNode) int {
	ret = 0
	dfs(root, 0)
	return  ret
}

func dfs(root *TreeNode, sum int)  {
	if root == nil {
		return
	}
	sum = sum << 1
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		ret = ret + sum
		return
	}

	dfs(root.Left, sum)
	dfs(root.Right, sum )

}

/*
var ret int

func sumRootToLeaf(root *TreeNode) int {
	ret = 0
	helper(root, 0)
	return ret
}
func helper(root *TreeNode, sum int) {
	if root == nil {
		return
	}
	sum = sum << 1
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		ret += sum
		return
	}
	if root.Left != nil {
		helper(root.Left, sum)
	}
	if root.Right != nil {
		helper(root.Right, sum)
	}
}
*/