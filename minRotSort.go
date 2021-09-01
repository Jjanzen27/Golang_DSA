/* Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

[4,5,6,7,0,1,2] if it was rotated 4 times.
[0,1,2,4,5,6,7] if it was rotated 7 times.
Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums of unique elements, return the minimum element of this array.

You must write an algorithm that runs in O(log n) time.



Example 1:

Input: nums = [3,4,5,1,2]
Output: 1
Explanation: The original array was [1,2,3,4,5] rotated 3 times.
Example 2:

Input: nums = [4,5,6,7,0,1,2]
Output: 0
Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.
Example 3:

Input: nums = [11,13,15,17]
Output: 11
Explanation: The original array was [11,13,15,17] and it was rotated 4 times.


Constraints:

n == nums.length
1 <= n <= 5000
-5000 <= nums[i] <= 5000
All the integers of nums are unique.
nums is sorted and rotated between 1 and n times.
*/

package main

import (
	"fmt"
)

// func findMin(nums []int) int {
// 	start, last := 0, len(nums)-1
// 	firstValue := nums[0]
// 	if nums[0] < nums[last] || last == 0 {
//     return nums[0]
// }

// 	for start < last {
// 		last = len(nums) - 1
// 		midpoint := last / 2
// 		if firstValue <= nums[midpoint] {
// 			nums = nums[midpoint+1 : last+1]
// 		} else {
// 			nums = nums[start : midpoint+1]
// 		}
// 	}
// 	return nums[0]
// }

func findMin(nums []int) int {

	low := 0
	high := len(nums) - 1

	for low < high {
		if nums[low] < nums[high] || low == high {
			break
		}

		mid := low + (high-low)/2

		if nums[mid] < nums[low] {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return nums[low]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{1, 2, 3}))
	fmt.Println(findMin([]int{3, 4, 5, 6, 7, 1, 2}))
	fmt.Println(findMin([]int{2, 3, 4, 5, 1}))
}
