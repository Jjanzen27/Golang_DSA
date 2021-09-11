/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 
 /**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 
 Mental model: We need to use a dummy head in case the first value
 of the linked list matches the target.
 
 There are two things we need to track: the current node and the previous node. This is because in order to delete a node, we need to reassign previousNode.nextNode to the node *after* the currentNode (effectively skipping over it -- it still exists out there in memory, but it's no longer attached to the list.)
 
 d -> 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6 ->
 p    c
 
 *** Check to see if the current node is equal to the target***
 Break condition: currentNode.Next will be nil
 
    It's not. Move both pointers to the right.
    
d -> 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6 ->
     p    c
 
    Repeat
    
d -> 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6 ->
          p    c
          
    It is!!
    
    So now we reassign previousNode.nextNode to currentNode.nextNode
    
d -> 1 -> 2 ->  -> 3 -> 4 -> 5 -> 6 ->
          p         c
    
    We reassign currentNode to currentNode.Next, but don't reassign 
    previousNode, in case there are adjacent target values.
    

 */
func removeElements(head *ListNode, val int) *ListNode {
    var previousNode *ListNode
    dummyNode := &ListNode{Val: -1, Next: head}
    previousNode = dummyNode
    currentNode := head
    
    for currentNode != nil {
        if currentNode.Val == val {
            previousNode.Next = currentNode.Next
        } else {
            previousNode = currentNode
        }
        currentNode = currentNode.Next
    }
    
    return dummyNode.Next
    
}
 
/*
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
    dummyNode := &ListNode{Val: -1, Next: head}
    previousNode = dummyNode
    currentNode := head
    
    for currentNode != nil {
        if currentNode.Val == val {
            previousNode.Next = currentNode.Next
        }
        previousNode = currentNode
        currentNode = currentNode.Next
    }
    
    return dummyNode.Next
    
}