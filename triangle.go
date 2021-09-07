// Given a triangle array, return the minimum path sum from top to bottom.

// For each step, you may move to an adjacent number of the row below.
// More formally, if you are on index i on the current row, you may move to either
// index i or index i + 1 on the next row.

// Example 1:

// Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
// Output: 11
// Explanation: The triangle looks like:
//   2
//   3 4
//  6 5 7
// 4 1 8 3
// The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined above).
// Example 2:

// Input: triangle = [[-10]]
// Output: -10

// Constraints:

// 1 <= triangle.length <= 200
// triangle[0].length == 1
// triangle[i].length == triangle[i - 1].length + 1
// -104 <= triangle[i][j] <= 104

// Follow up: Could you do this using only O(n) extra space, where n is the
// total number of rows in the triangle?

/*

First attempt: doesn't use dp, works in some cases (starts at top and chooses lesser
value until bottom). Doesn't work if the minimum sum is on a path with a higher value
at any point.

*/

// package main

// import (
// 	"fmt"
// )

// func minimumTotal(triangle [][]int) int {
// 	sum := 0
// 	return triangleHelper(triangle, sum, 0, 0)
// }

// func triangleHelper(triangle [][]int, sum int, index int, level int) int {
// 	if level >= len(triangle) {
// 		return sum
// 	}

// 	if level == 0 {
// 		index = 0
// 	} else if triangle[level][index] > triangle[level][index+1] {
// 		fmt.Println(triangle[level][index], triangle[level][index+1])
// 		index += 1
// 	}

// 	sum += triangle[level][index]

// 	return triangleHelper(triangle, sum, index, level+1)
// }

// base case: when we get to bottom of triangle

//    2
//   3 4
//  6 5 7
// 4 1 8 3

// on the first level, there's only one option
// on the second level, there are two options: i == 0 or 1 [0][0] or [0][1]
// on the third level, we can go [0][0][0], [0][0][1], [0][1][1] or [0][1][2]
// on the fourth level [0][0][0][0], [0][0][0][1], [0][0][1][1], [0][0][1][2],
//                     [0][1][1][1], [0][1][1][2], [0][1][2][2], [0][1][2][3]

// map:
// [0] :2
// [0, 0] : 5           add together the result of map[0] and triangle[1][0]
// [0, 1] : 6           add together the result of map[0] and triangle[1][1]
// [0, 0, 0] : 11       add together map[0, 0] and triangle[2][0]
// }
package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	dp := make(map[[2]int]int)

	for row := len(triangle) - 1; row >= 0; row-- {
		for col := len(triangle[row]) - 1; col >= 0; col-- {
			if row == len(triangle)-1 {
				dp[[2]int{row, col}] = triangle[row][col]
			} else {
				// min(sum of their value + square below and sum of value of square below and right)
				dp[[2]int{row, col}] = min(dp[[2]int{row + 1, col}], dp[[2]int{row + 1, col + 1}]) + triangle[row][col]
			}
		}

	}
	return dp[[2]int{0, 0}]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTotal([][]int{{1}, {2, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
	fmt.Println(minimumTotal([][]int{{-1}, {2, 3}, {1, -1, -3}}))
}
