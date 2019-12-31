package other

/*
二叉搜索树的特点是，左节点比根节点小，右节点比根节点大
对于这道题来说，由于要取 [L, R] 范围的树
如果这个节点比 R 大，那么他的右节点更大，丢掉，去找他的左节点
如果这个节点比 L 小，那么他的左节点更小，丢掉，去找他的右节点
属于这个范围之间的，继续遍历即可
*/
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