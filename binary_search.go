/* Binary Search
Given an array of integers nums which is sorted in ascending order, and an integer target,
write a function to search target in nums. If target exists, then return its index.
Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.

Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is

Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1

Constraints:
1 <= nums.length <= 10^4
-104 < nums[i], target < 104
All the integers in nums are unique.
nums is sorted in ascending order.

-------------------------------------------------------------

First solution (O(N)):

package main

import (
	"fmt"
)

func binarySearch(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		} else if v > target {
			return -1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{-1,0,3,5,9,12}, 2))
}


*/

package main

import "fmt"

func binarySearch(nums []int, target int) int {
	lowerBound := 0
	upperBound := len(nums) - 1

	for lowerBound <= upperBound {
		midpoint := (upperBound + lowerBound) / 2
		if nums[midpoint] == target {
			return midpoint
		} else if nums[midpoint] > target {
			upperBound = midpoint - 1
		} else if nums[midpoint] < target {
			lowerBound = midpoint + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(binarySearch([]int{-1, 0, 3, 5, 9, 12}, 2))
}
