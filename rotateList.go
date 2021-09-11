// Given the head of a linked list, rotate the list to the right by k places.

// Example 1:

// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]
// Example 2:

// Input: head = [0,1,2], k = 4
// Output: [2,0,1]

// Constraints:

// The number of nodes in the list is in the range [0, 500].
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 109

/*
Algorithm:

Brute force:
Traverse LL to find the last node (keep track of second-to-last too)
Set last node's nextNode to current head
Set second-to-last node's nextNode to nil
Repeat k times
(Have to keep searching through list for last node because can't
move backwards through list)

A better solution:
Find the node at position k, the one before and the one at the end
Set the one before k's next node to nil
Set the end node's nextNode to head

Edge cases: If k is greater than the length of the list, the first solution
would still work, but the second one would break

- Start with three pointers at the head and a variable to track length
- Send the first tracker out k steps, incrementing the length tracker
    - If the tracker hits `nil`, we know our length is less than k
    - Find the remainder of k/len and set k to that
- Send out k and k-1 pointers k nodes (don't actually need k pointer)
- Set the node at k-1's nextNode to nil
- Set the node at the end's pointer to the head

*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	pointer := 1
	currentNode := head
	targetNode := head

	for currentNode.Next != nil && pointer <= k {
		currentNode = currentNode.Next
		pointer++
	}

	if pointer < k {
		k = k % pointer
		pointer = pointer - k
	}

	for currentNode.Next != nil {
		currentNode = currentNode.Next
		targetNode = targetNode.Next
	}
	currentNode.Next = head
	targetNode.Next = nil

	return currentNode
}

func main() {
	// 	node8 := &ListNode{Val: 8}
	// 	node7 := &ListNode{Val: 7, Next: node8}
	// 	node6 := &ListNode{Val: 6, Next: node7}
	// 	node5 := &ListNode{Val: 5, Next: node6}
	// 	node4 := &ListNode{Val: 4, Next: node5}
	node3 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2, Next: node3}
	node1 := &ListNode{Val: 1, Next: node2}

	// fmt.Println(rotateRight(node1, 1))
	fmt.Println(rotateRight(node1, 4))
	//fmt.Println(node3.Next.Next.Next.Next)
}
