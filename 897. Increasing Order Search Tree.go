package main


  //Definition for a binary tree node.
  //type TreeNodee struct {
  //    Val int
  //    Left *TreeNode
  //    Right *TreeNode
  //}
func increasingBST(root *TreeNode) *TreeNode {
	inOrder(root)
	return result
}

var result *TreeNode
var last *TreeNode

func inOrder(root *TreeNode) {
	if root != nil {
		inOrder(root.Left)
		temp := TreeNode{Val:root.Val}
		if last != nil {
			last.Right = &temp
			last = &temp
		} else  {
			result = &temp
			last = &temp
		}

		inOrder(root.Right)
	}
}