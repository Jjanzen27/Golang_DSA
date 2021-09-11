/*
Given an array nums of distinct integers, return all the possible
permutations. You can return the answer in any order.


Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]]


Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/

package main

import (
	"fmt"
)

func permutations(nums []int) (result [][]int) {
	current := []int{}
	for i := 0; i < len(nums); i++ {
		current = append(current, nums[i])
		for j := 0; j < len(nums); j++ {
			if i != j {
				current = append(current, nums[j])
			}
		}
		result = append(result, current)
		current = []int{}
	}
	return
}

func main() {
	fmt.Println(permutations([]int{1, 2, 3}))
}
