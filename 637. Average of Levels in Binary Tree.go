package main

import (
	"fmt"
)

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val float
*     Left *TreeNode
*     Right *TreeNode
* }
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {

	root := &TreeNode{Val:4}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:6}
	root.Left.Left = &TreeNode{Val:3}
	root.Left.Right = &TreeNode{Val:5}
	root.Right.Right = &TreeNode{Val:7}

	fmt.Println(averageOfLevels(root))
}

func averageOfLevels(root *TreeNode) []float64 {
	return bfs1(root)
}
/*
广度优先遍历, 一层一层来
*/
func bfs1(root *TreeNode) []float64  {
	res := make([]float64, 0)

	queue := make([]*TreeNode, 0)
	if root == nil {
		return res
	}

	queue = append(queue, root)

	tempQueue := make([]*TreeNode, 0)
	currentSum := 0
	count := 0
	for ; len(queue) > 0;   {
		node := queue[0]
		currentSum += node.Val
		count += 1
		queue = queue[1:]
		if node.Left != nil {
			tempQueue = append(tempQueue, node.Left)
		}
		if node.Right != nil {
			tempQueue = append(tempQueue, node.Right)
		}

		if len(queue) == 0 {
			res = append(res, float64(currentSum) / float64(count))
			count = 0
			currentSum = 0
			queue = tempQueue
			tempQueue = make([]*TreeNode, 0)
		}
	}
	return res
}