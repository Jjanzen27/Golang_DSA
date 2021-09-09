/*
Given the root of a binary tree, return the preorder traversal of its nodes' values.

Example 1:


Input: root = [1,null,2,3]
Output: [1,2,3]
Example 2:

Input: root = []
Output: []
Example 3:

Input: root = [1]
Output: [1]
Example 4:


Input: root = [1,2]
Output: [1,2]
Example 5:


Input: root = [1,null,2]
Output: [1,2]


Constraints:

The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100


Follow up: Recursive solution is trivial, could you do it iteratively?

With pop method:

func preorderTraversal(root *TreeNode) (result []int) {
    stack := []*TreeNode{root}

    for len(stack) > 0 {
        node := pop(&stack)

        if node != nil {
            result = append(result, node.Val)
            stack = append(stack, node.Right)
            stack = append(stack, node.Left)
        }
    }

    return
}

func pop(slice *[]*TreeNode) *TreeNode {
    node := (*slice)[len(*slice)-1]
    *slice = (*slice)[:len(*slice)-1]
    return node
}
*/

package main

import (
	"fmt"
)

func preorderTraversal(root *TreeNode) (result []int) {
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			result = append(result, node.Val)
			stack = append(stack, node.Right)
			stack = append(stack, node.Left)
		}
	}

	return
}

func main() {

}
