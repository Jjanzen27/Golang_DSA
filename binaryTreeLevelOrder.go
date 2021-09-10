// Given the root of a binary tree, return the level order traversal of its nodes'
// values. (i.e., from left to right, level by level).

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]
// Example 2:

// Input: root = [1]
// Output: [[1]]
// Example 3:

// Input: root = []
// Output: []

// Constraints:

// The number of nodes in the tree is in the range [0, 2000].
// -1000 <= Node.val <= 1000

package main

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func levelOrder(root *TreeNode) [][]int {
// 	result := [][]int{}
// 	queue := []*TreeNode{}
// 	currentLevel := []int{}

// 	queue = append(queue, root)

// 	for len(queue) > 0 {
// 		currentLevel = append(currentLevel, queue[0].Val)
// 		if queue[0]

// 		fmt.Println(currentLevel, result)
// 	}

// 	return result
// }

func levelOrder(root *TreeNode) []int {
	result := []int{}

	return helper(root, result)

}

func helper(node *TreeNode, res []int) []int {
	if node != nil {
		res = append(res, node.Val)
		if node.Left != nil {
			res = helper(node.Left, res)
		}
		if node.Right != nil {
			res = helper(node.Right, res)
		}
	}
	return res
}

func main() {
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 3, Left: node6}
	node2 := &TreeNode{Val: 2, Left: node4, Right: node5}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}

	fmt.Println(levelOrder(node1))
}
