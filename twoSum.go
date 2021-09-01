/*
Two Sum

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n2)
time complexity?

Input: An array and a target number
Output: Two indices, the locations of two numbers that add up to target

Naive approach (O(N^2)): create two nested loops and check if every
value added to every other value equals the target

Better approach: use a built-in sorting algorithm to sort data in
o(N log N) time then start double loop with an early exit if two
sums are greater than the target

Optimized approach: use a hash map to track which values we've seen,
and check if target - current is in the map already

*/

package main

import (
	"fmt"
)

func twoSum(array []int, target int) []int {
	m := make(map[int]int)

	for i, v := range array {
		if _, ok := m[target-v]; ok {
			return []int{m[target-v], i}
		} else {
			m[v] = i
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
