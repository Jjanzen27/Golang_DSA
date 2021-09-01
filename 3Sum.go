/*
15. 3Sum
Given an integer array nums, return all the triplets [nums[i], nums[j],
nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] +
nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Example 2:

Input: nums = []
Output: []
Example 3:

Input: nums = [0]
Output: []


Constraints:

0 <= nums.length <= 3000
-105 <= nums[i] <= 105

func threeSum(nums []int) (triplets [][]int) {
	start, end := nums[0], len(nums)-1
	var tripletPointer int

	for start < end {
		twoSum = start + end

		if twoSum < 0 {
			tripletPointer = end - 1
		} else {
			tripletPointer = start + 1
		}

		triSum := twoSum + tripletPointer

		if triSum == 0 {
			triplets = append(triplets, []int{nums[start], nums[end], nums[tripletPointer]})
		} else if triSum*twoSum > 0 { // both have the same sign
			start += 1
			end = len(nums) - 1
		}
	}

}


_________________________________
Input: An array of numbers (positive and negative)
Output: A nested array: all sets of three numbers that add up to 0
Rules: No duplicate sets of triplets and no repeating elements (indexes must be different)

Examples:

threeSum([]int{-1,0,1,2,-1,-4}) // [[-1,-1,2],[-1,0,1]]
threeSum([]int{}) //[]
threeSum([]int{0}) //[]

Ideas:
Sort input array

[-4, -1, -1, 0, 1, 2]

Start with nums[i] and nums[len(nums)-1]
Add the two together (-4 + 2 = -2)
If the number is negative, move end pointer to the left (now pointing at 1)
If the number is positive, move start pointer to the right
Check to see if the sum of the three numbers equals 0 (-4 + 2 + 1) = -1
    If it's 0, add the triplet to return
    If it's not, but the sum has switched signs (become negative or positive), move same pointer and check again
    If it's not, but the sum is the same sign, the current start isn't part of a triplet.
        Increment start by 1 and reset end to the length of the array


Return when pointers equal each other


Algo:
Point at opposite ends of the array
Add those two numbers together

If current sum negative, pointer goes at end - 1 : start + 1

Start loop to iterate while pointer is greater than/less than 0
    and sum of s+e+pointer >= 0

Add currentSum to pointer number
    Check to see if 0 -> if it is,


*/

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func threeSum(nums []int) (trips [][]int) {
// 	sort.Ints(nums)

// 	for start, end := 0, len(nums)-1; nums[start] <= 0; {
// 		currentSum := nums[start] + nums[end]
// 		var pointer int
// 		var sign string

// 		fmt.Println(currentSum)

// 		if currentSum >= 0 {
// 			pointer = start + 1
// 			sign = "negative"
// 		} else {
// 			pointer = end - 1
// 			sign = "positive"
// 		}

// 		fmt.Println(sign, pointer, start, end)

// 		if sign == "negative" {
// 			for nums[pointer] < 0 && currentSum+nums[pointer] >= 0 {
// 				fmt.Println(currentSum, nums[pointer], currentSum+nums[pointer], pointer)
// 				if currentSum+nums[pointer] == 0 {
// 					trips = append(trips, []int{nums[start], nums[end], nums[pointer]})
// 				}
// 				pointer += 1
// 			}

// 		} else {
// 			for nums[pointer] > 0 && currentSum+nums[pointer] <= 0 {
// 				fmt.Println(currentSum, nums[pointer], currentSum+nums[pointer], pointer)
// 				if currentSum+nums[pointer] == 0 {
// 					trips = append(trips, []int{nums[start], nums[end], nums[pointer]})
// 				}
// 				pointer -= 1
// 			}
// 		}

// 		start += 1

// 		fmt.Println(trips)
// 	}

// 	return
// }

// func main() {
// 	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
// 	fmt.Println(threeSum([]int{-8, -6, -3, -3, -1, 0, 1, 1, 2, 5, 6}))
// }

/*
Split into two arrays, negative numbers and positive numbers
Call twoSum on negative array, passing each number of positive array as target
If the three add up to 0, create an array of them and add to triplets
Repeat reverse


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
	return nil
}

func filter(ss []int, test func(int) bool) (ret []int) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func isNegative(num int) bool {
	return num <= 0
}

func isPositive(num int) bool {
	return num > 0
}

func isZero(num int) bool {
	return num == 0
}

func contatins(array [][]int, slice []int ) bool {
    for i, v := range array {
        if 
    }
}

func threeSum(nums []int) (trips [][]int) {
	negatives := filter(nums, isNegative)
	positives := filter(nums, isPositive)
	zeroes := filter(nums, isZero)

	if len(zeroes) >= 3 {
		for i := 0; i < len(zeroes)-1; i++ {
			trips = append(trips, []int{0, 0, 0})
		}
	}

	for i := 0; i < len(negatives); i++ {
		target := -negatives[i]

		if twoSum(positives, target) != nil {
			indices := twoSum(positives, target)

			trips = append(trips, []int{positives[indices[0]], positives[indices[1]], negatives[i]})
		}
	}

	for i := 0; i < len(positives); i++ {
		target := -positives[i]

		if twoSum(negatives, target) != nil {
			indices := twoSum(negatives, target)

			trips = append(trips, []int{negatives[indices[0]], negatives[indices[1]], positives[i]})
		}
	}

	return removeDuplicates(trips)
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-8, -6, -3, -3, -1, 0, 1, 1, 2, 5, 6}))
}
