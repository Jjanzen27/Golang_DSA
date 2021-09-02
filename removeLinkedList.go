/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
    Val int
    Next *ListNode
}

// func newNode(Val int) *node {
//     n := node{Val: val}
//     return &node
// }
 
func removeElements(head *ListNode, val int) *ListNode {
    var previousNode *ListNode
    dummyNode := &ListNode{Val: -1, Next: &head}
    previousNode = dummyNode
    currentNode := head
    
    for currentNode != nil {
        if currentNode.Val == val {
            previousNode.Next = currentNode.Next
        }
        previousNode = currentNode
        currentNode = currentNode.Next
    }
    
    return dummyNode.next
    
}