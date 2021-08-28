/*
Find First and Last Position of Element in Sorted Array

Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
Example 3:

Input: nums = [], target = 0
Output: [-1,-1]


Constraints:

0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums is a non-decreasing array.
-109 <= target <= 109

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

func searchRange(nums []int, target int) []int {
	searchIndex := binarySearch(nums, target)

	if searchIndex == -1 {
		return []int{-1, -1}
	}

	x, y := searchIndex, searchIndex
	for x < len(nums) && nums[x] == target {
		x += 1
	}

	for y >= 0 && nums[y] == target {
		y -= 1
	}

	return []int{y + 1, x - 1}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8)) // [3, 4]
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6)) // [-1,-1]
	fmt.Println(searchRange([]int{}, 0))                  // [-1, -1]
}
