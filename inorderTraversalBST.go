/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.

Example 1:


Input: root = [1,null,2,3]
Output: [1,3,2]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]
Example 4:


Input: root = [1,2]
Output: [2,1]
Example 5:


Input: root = [1,null,2]
Output: [1,2]


Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100


Follow up: Recursive solution is trivial, could you do it iteratively?
*/

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
func inorderTraversal(root *TreeNode) (nodes []int) {
	if root == nil {
		return []int{}
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	left = append(left, root.Val)
	left = append(left, right...)
	return left
}

func main() {

}

// stack := []int{}
// return helper(root, &stack)

// func helper() {

// }
