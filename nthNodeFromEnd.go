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
 
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     s := head
//     e := head
    
//     for i := 1; i < n+1; i++ {
//         if s.Next == nil {
//             break
//         }
//         s = s.Next
//     }
    
//     for s != nil {
//         e = e.Next
//         s = s.Next
//     }
    
//     return head
// }

/*
Problem: Given the head of a linked list, remove the nth node from the end of the list and return its head.

[1] -> [2] -> [3] -> [4] -> [5]     n = 2
                      ^ 

Mental model: We can find the correct node to delete by using two pointers. 
Both should start at the same position.
Move one of the pointers n spaces.

[1] -> [2] -> [3] -> [4] -> [5]     n = 2
p1
p2                        

Then, increment both pointers at the same rate until the further one's next node is `nil`.

[1] -> [2] -> [3] -> [4] -> [5]     n = 2
              p1
                             p2  

At that point, delete the relevant node by assigning the slower pointer's next node to it's next next.

The tricky part is handling the cases where we need to delete the first or only node (if n == len(ll))

[1]         n = 1
p1
p2

[1, 2]     n = 2
p1
       p2

Solution: add a condition that "chops off the head", if the faster node encounters `nil` 


*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    fast := head
    slow := head
    
    for i := 0; i < n+1; i++ {
        if fast == nil {
            return head.Next
        }
        fast = fast.Next
    }
    
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }

    slow.Next = slow.Next.Next
    
    return head
}


/*

Solution using a dummy head:

*/

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  dummyNode := &ListNode{0, head}
  slow := dummyNode  
  fast := dummyNode  
  
  for i := 0; i <= n; i += 1 {
    if fast == nil {
      break
    }
    fast = fast.Next
  }
  
  for fast != nil {
    slow = slow.Next
    fast = fast.Next
  }
  
  if slow.Next != nil {
    slow.Next = slow.Next.Next
  }
  
  return dummyNode.Next
}