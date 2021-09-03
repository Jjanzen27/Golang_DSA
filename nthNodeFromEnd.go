// Given the head of a linked list, remove the nth node from the end of the list and return its head.

 

// Example 1:


// Input: head = [1,2,3,4,5], n = 2
// Output: [1,2,3,5]
// Example 2:

// Input: head = [1], n = 1
// Output: []
// Example 3:

// Input: head = [1,2], n = 1
// Output: [1]
 

// Constraints:

// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
 

// Follow up: Could you do this in one pass?

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     s := head
//     e := head
    
//     for i := 0; i < n; i++ {
//         s = s.next
//     }
    
//     for e != nil {
//         e = e.next
//         s = s.next
//     }
    
//     return e
// }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
 // Neither solution works... despite being the same as the one in Cracking 
 // the coding interview. The issue is with linked lists 2 elements or shorter
 
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    s := head
    e := head
    
    for i := 1; i < n+1; i++ {
        if s.Next == nil {
            return nil
        }
        s = s.Next
    }
    
    for s != nil {
        e = e.Next
        s = s.Next
    }
    
    return head
}