// Merge two sorted linked lists and return it as a new list.
// The new list should be made by splicing together the nodes
// of the first two lists.

/*

 */

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) {
	dummyNode := &ListNode{Val: -1, Next: nil}
	currentNode := dummyNode

	for l1.Next != nil && l2.Next != nil {
		if l1.val > l2.val {
			currentNode.Next = l1
			l1 = l1.Next
		} else {
			currentNode.Next = l2
			l2 = l2.Next
		}
		currentNode = currentNode.Next
	}

	if l1.Next != nil {
		currentNode.next = l1
	}
	if l2.Next != nil {
		currentNode.next = l2
	}

	return dummyHead.Next
}

func main() {

}
