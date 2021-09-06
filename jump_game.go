// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

// Return true if you can reach the last index, or false otherwise.

// Example 1:

// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:

// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

// Constraints:

// 1 <= nums.length <= 104
// 0 <= nums[i] <= 105

package main

import (
	"fmt"
)

// func canJump(nums []int) bool {
// 	availSteps := nums[0]

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] > availSteps {
// 			availSteps = nums[i]
// 		} else if availSteps <= 0 && i != len(nums)-1 {
// 			return false
// 		}
// 		availSteps--
// 	}
// 	return true
// }

func canJump(nums []int) bool {
	availSteps := nums[0]
	return helper(nums, availSteps)
}

func helper(nums []int, availSteps int) bool {
	//fmt.Println(nums, availSteps)
	if len(nums) == 1 {
		return true
	}

	if nums[0] > availSteps {
		availSteps = nums[0]
	}

	if availSteps == 0 {
		return false
	}

	return helper(nums[1:], availSteps-1)
}

func main() {
	fmt.Println(canJump([]int{1, 2, 3}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{1, 1, 0, 1}))
}
