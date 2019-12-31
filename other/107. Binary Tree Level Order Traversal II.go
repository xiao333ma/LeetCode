package other

import "fmt"

///**
// * Definition for a binary tree node.
// * type TreeNode struct {
// *     Val int
// *     Left *TreeNode
// *     Right *TreeNode
// * }
// */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func main() {

	root := &TreeNode{Val:3}
	root.Left = &TreeNode{Val:9}
	root.Right = &TreeNode{Val:20}
	root.Right.Left = &TreeNode{Val:15}
	root.Right.Right = &TreeNode{Val:7}

	fmt.Println(levelOrderBottom(root))
}
func levelOrderBottom(root *TreeNode) [][]int {

	return bfs(root)
}

func bfs(root *TreeNode) [][]int  {

	res := make([][]int, 0)

	if root == nil {
		return res
	}

	q := make([]*TreeNode, 0)

	currentLayer := make([]*TreeNode, 0)
	currentVale := make([]int, 0)
	q = append(q, root)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			currentLayer = append(currentLayer, node.Left)
		}

		if node.Right != nil {
			currentLayer = append(currentLayer, node.Right)
		}
		currentVale = append(currentVale, node.Val)

		if len(q) == 0 {
			q = currentLayer
			currentLayer = make([]*TreeNode, 0)
			res = append(res, currentVale)
			currentVale = make([]int, 0)
		}
	}
	r := make([][]int, 0)
	for i := len(res) - 1; i >= 0; i -- {
		r = append(r, res[i])
	}

	return r
}