package main


  //Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }

func increasingBST(root *TreeNode) *TreeNode {
	inOrder(root)

	tree := new(TreeNode)

	for _, value := range array {

	}
}

var array = make([]int, 0);

func inOrder(root *TreeNode) {
	if root != nil {
		inOrder(root.Left)
		array = append(array, root.Val)
		inOrder(root.Right)
	}
}