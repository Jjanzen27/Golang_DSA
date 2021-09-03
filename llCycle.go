// Given head, the head of a linked list, determine if the linked list has a cycle in it.

// There is a cycle in a linked list if there is some node in 
// the list that can be reached again by continuously following 
// the next pointer. Internally, pos is used to denote the index 
// of the node that tail's next pointer is connected to. 

// Note that pos is not passed as a parameter.

// Return true if there is a cycle in the linked list. Otherwise, return false.

 

// Example 1:

// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to 
// the 1st node (0-indexed).

// Example 2:

// Input: head = [1,2], pos = 0
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
// Example 3:


// Input: head = [1], pos = -1
// Output: false
// Explanation: There is no cycle in the linked list.
 

// Constraints:

// The number of the nodes in the list is in the range [0, 104].
// -105 <= Node.val <= 105
// pos is -1 or a valid index in the linked-list.
 
// Follow up: Can you solve it using O(1) (i.e. constant) memory?

/*

Problem: Check if there's a cycle in a linked list.

Explanation: There's a cycle in a list if there is some node in 
the list that can be reached again by continuously following 
the next pointer.

[1] -> [2] -> [3] -> [4] -> [5]
        ^                    |  
        |                    |
        |____________________|

Mental model: Follow the next pointer until we either reach a node 
we've visited before or nil.

    --> If we hit 'nil', there isn't a cycle: return false

Question: How can we tell if we've already visited a node?

Option 1: Use a hash map to keep track of the nodes we've visited
as we iterate through the list. If we find a node that's already
in the hash map, return true.

Option 2: Use fast and slow pointers. The slow pointer moves one
node at a time. The fast pointer moves two nodes at a time. 

If the fast pointer reaches nil, there are no cycles in our list.

If the fast pointer catches up to the slow pointer, we found a cycle!

fast pointer = f
slow pointer = s

 s
 f
[1] -> [2] -> [3] -> [4] -> [5]
        ^                    |  
        |                    |
        |____________________|
        
        
        s
               f
[1] -> [2] -> [3] -> [4] -> [5]
        ^                    |  
        |                    |
        |____________________|


               s
                             f
[1] -> [2] -> [3] -> [4] -> [5]
        ^                    |  
        |                    |
        |____________________|

                      s
               f
[1] -> [2] -> [3] -> [4] -> [5]
        ^                    |  
        |                    |
        |____________________|
        
                             s
                             f
[1] -> [2] -> [3] -> [4] -> [5]
        ^                    |  
        |                    |
        |____________________|
        
The pointers are pointing at the same node, so we have a loop!

*/


func hasCycle(head *ListNode) bool {
    slow := head
    fast := head
    
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            return true
        }
    }
    
    return false
}