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

// func canJump(nums []int) bool {
// 	availSteps := nums[0]
// 	return helper(nums, availSteps)
// }

// func helper(nums []int, availSteps int) bool {
// 	//fmt.Println(nums, availSteps)
// 	if len(nums) == 1 {
// 		return true
// 	}

// 	if nums[0] > availSteps {
// 		availSteps = nums[0]
// 	}

// 	if availSteps == 0 {
// 		return false
// 	}

// 	return helper(nums[1:], availSteps-1)
// }

// Yujohn's solution

// func canJump(nums []int) bool {
// 	return helper(nums, len(nums)-2, len(nums)-1)
// }

// func helper(nums []int, index int, targetIndex int) bool {
// 	if index <= 0 {
// 		return nums[0] >= targetIndex
// 	}

// 	if index+nums[index] >= targetIndex {
// 		return helper(nums, index-1, index)
// 	}

// 	return helper(nums, index-1, targetIndex)
// }

// func canJump(nums []int) bool {
//     n = len(nums)
//     steps := [n]bool

//     nums[0] = true

//     for i, _ := range(1, nums) {
//         temp := false

//         for j, _ := range(i-1, -1, -1) {
//             if nums[j] > i - j && steps[j] == true {
//                 steps[i] == true
//                 break
//             }
//         }
//     }
//     return
// }

func canJump(nums []int) bool {
	if len(nums) < 1 {
		return true
	}
	memo := make(map[int]bool)
	return jump(0, nums, memo)

}

func jump(pos int, nums []int, memo map[int]bool) bool {
	if pos == len(nums)-1 {
		return true
	}
	val, ok := memo[pos]

	if ok {
		return val
	}

	if pos >= len(nums) || nums[pos] == 0 {
		return false
	}

	for i := pos + nums[pos]; i >= pos+1; i-- {
		if jump(i, nums, memo) {
			memo[pos] = true
			return true
		}
	}
	memo[pos] = false

	return false

}

func main() {
	fmt.Println(canJump([]int{1, 2, 3}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{0}))
	fmt.Println(canJump([]int{1, 1, 0, 1}))
}
