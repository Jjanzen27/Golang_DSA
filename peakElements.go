// A peak element is an element that is strictly greater than its neighbors.

// Given an integer array nums, find a peak element, and return its index.
// If the array contains multiple peaks, return the index to any of the peaks.

// You may imagine that nums[-1] = nums[n] = -∞.

// You must write an algorithm that runs in O(log n) time.

// Example 1:

// Input: nums = [1,2,3,1]
// Output: 2
// Explanation: 3 is a peak element and your function should return the index number 2.
// Example 2:

// Input: nums = [1,2,1,3,5,6,4]
// Output: 5
// Explanation: Your function can return either index number 1 where the peak element is 2,
// or index number 5 where the peak element is 6.

// Constraints:

// 1 <= nums.length <= 1000
// -231 <= nums[i] <= 231 - 1
// nums[i] != nums[i + 1] for all valid i.

/*
- Grab element at midpoint/pivot
- Check to see if elements on either side are higher or lower than it
- If both are lower, return it
- Otherwise, move toward the higher one
- Check to see if the element on the other side of that is lower
- Repeat til find peak or edge
*/

package main

import (
	"fmt"
)

func findPeakElement(nums []int) int {

	start, end := 0, len(nums)-1
	mid := start + ((end - start) / 2)
	previous := mid - 1
	next := mid + 1

	for start < end {

		if previous < start && next > end {
			return 0
		}

		if previous < start {
			if nums[mid] > nums[next] {
				return mid
			} else {
				return next
			}
		}

		if next > end {
			if nums[mid] > nums[previous] {
				return mid
			} else {
				return previous
			}
		}

		if nums[mid] > nums[previous] && nums[mid] > nums[next] {
			return mid
		} else if nums[previous] > nums[next] {
			mid = previous
			next = mid
			previous -= 1
		} else {
			previous = mid
			mid = next
			next += 1
		}
	}
	return 0
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2}))
}
