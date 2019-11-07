package main


func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val > R { // 比 R 大，只要左边的
		return trimBST(root.Left, L, R)
	}
	if root.Val < L { // 比 L 小，只要右边的
		return trimBST(root.Right, L, R)
	}
	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return  root
}