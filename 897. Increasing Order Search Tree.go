package main


  //Definition for a binary tree node.
  //type TreeNodee struct {
  //    Val int
  //    Left *TreeNode
  //    Right *TreeNode
  //}
func increasingBST(root *TreeNode) *TreeNode {
	inOrder(root)
	return result.Right
}

var result *TreeNode = &TreeNode{}
var last *TreeNode = result

func inOrder(root *TreeNode) {
	if root != nil {
		inOrder(root.Left)
		temp := &TreeNode{Val:root.Val}
		temp.Left = nil
		last.Right = temp
		last = temp
		inOrder(root.Right)
	}
}