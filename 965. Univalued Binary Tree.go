package main

/*
A binary tree is univalued if every node in the tree has the same value.

Return true if and only if the given tree is univalued.



Example 1:


Input: [1,1,1,1,1,null,1]
Output: true
Example 2:


Input: [2,2,2,5,2]
Output: false


Note:

The number of nodes in the given tree will be in the range [1, 100].
Each node's value will be an integer in the range [0, 99].

解题思路：

迭代的，取 根节点 为基数，遍历树，找到和节点不一样的值，则返回

*/


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 var num = -1

 // 迭代
func isUnivalTree(root *TreeNode) bool {
	arr := make([]*TreeNode, 0)
	baseValue := root.Val
	if root.Left != nil {
		arr = append(arr, root.Left)
	}
	if root.Right != nil {
		arr = append(arr, root.Right)
	}
	for ; len(arr) >= 1;  {
		val := arr[0].Val
		if val != baseValue {
			return false
		}else {
			if arr[0].Left != nil {
				arr = append(arr, arr[0].Left)
			}
			if arr[0].Right != nil {
				arr = append(arr, arr[0].Right)
			}
			if len(arr) >= 2 {
				arr = arr[1:]
			} else {
				return true
			}
		}
	}
	return true
}