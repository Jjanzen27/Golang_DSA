// Divide and Conquer Approach

// Given an integer array nums, find the contiguous subarray
// (containing at least one number) which has the largest sum and
// return its sum.

// A subarray is a contiguous part of an array.

// Example 1:

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
// Example 2:

// Input: nums = [1]
// Output: 1
// Example 3:

// Input: nums = [5,4,-1,7,8]
// Output: 23

// Constraints:

// 1 <= nums.length <= 3 * 104
// -105 <= nums[i] <= 105

// Follow up: If you have figured out the O(n) solution, try coding
// another solution using the divide and conquer approach, which is
// more subtle.

// Divide the array in half

package main

import "fmt"

func maxSubArray(nums []int) int {
	right := len(nums) - 1
	return findMaxSubArray(nums, 0, right)
}

func findMaxSubArray(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}

	mid := (left + right) / 2
	leftMax := findMaxSubArray(nums, left, mid)
	rightMax := findMaxSubArray(nums, mid+1, right)
	crossMax := findMaxCross(nums, left, right, mid)

	return max(leftMax, rightMax, crossMax)
}

func findMaxCross(nums []int, left int, right int, mid int) int {
	leftSum := -999999
	rightSum := -999999
	sum := 0

	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	sum = 0

	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}

	return leftSum + rightSum
}

func max(values ...int) int {
	maxVal := values[0]

	for i := 0; i < len(values); i++ {
		if values[i] > maxVal {
			maxVal = values[i]
		}
	}
	return maxVal
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
