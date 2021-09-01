/*
344. Reverse String
Write a function that reverses a string. The input string is given as
an array of characters s.

Example 1:

Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:

Input: s = ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]


Constraints:

1 <= s.length <= 105
s[i] is a printable ascii character.


Follow up: Do not allocate extra space for another array. You must do this
by modifying the input array in-place with O(1) extra memory.

=================================================================================
Input: an array of bytes
Output: the same array of bytes, but in reverse order

Rules: mutate the array in place

Example: reverseString([]byte{'H', 'a', 'n', 'n', 'a', 'h'})  // ['h', 'a', 'n', 'n', 'a', 'H']










=================================================================================

*/

package main

import (
	"fmt"
)

func reverseString(str []byte) {
	for start, end := 0, len(str)-1; start < end; start++ {
		str[start], str[end] = str[end], str[start]
		end--
	}
}

func main() {
	rev := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(rev)
	fmt.Println(rev)
}
