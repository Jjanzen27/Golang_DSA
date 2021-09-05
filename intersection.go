/*
Given two integer arrays nums1 and nums2, return an array of their intersection. Each element in the result must be unique and you may return the result in any order.



Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4,5]
Output: [9,4]
Explanation: [4,9] is also accepted.


Constraints:

1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000

Intersection: an element that occurs in both arrays, no matter where.
If elements occur multiple times, only select one.

Brute force:

nums1 = [4,9,5], nums2 = [9,4,9,8,4,5]
             ^
                                  ^

Create a hash map and a new array. Space O(N), Time O(N + M)

Problem: find intersection of 2 arrs
Input: two arrays of integers
Output: a new unique array containing integers that exist in both arrs

Datastruture: Hash + slice
Algorithm:

- Iterate through the first array and create a map of all elements that exist 5:true, 6:true

- Iterate through second array, if an element in the map == true, add that element to the new array and change map value to false

- Return the new array

*/

/*

 */

package main

import (
	"fmt"
)

func intersection(arr1 []int, arr2 []int) (intersection []int) {
	m := make(map[int]bool)

	for _, v := range arr1 {
		m[v] = true
	}

	for _, v := range arr2 {
		if m[v] == true {
			intersection = append(intersection, v)
			m[v] = false
		}
	}

	return
}

func main() {
	fmt.Println(intersection([]int{9, 4, 9, 8, 4}, []int{}))
}
