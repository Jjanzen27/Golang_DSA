/*
Given a m x n grid filled with non-negative numbers, find a path from top left
to bottom right, which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example 1:

Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
Example 2:

Input: grid = [[1,2,3],[4,5,6]]
Output: 12


Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 100

*/

/*
Very similar to triangle problem mashed with uniquePaths
Can only move down or right, find minimum sum

Start at top left

Base case: reach square at bottom right

Input: grid = [[1,3,1],[1,5,1],[4,2,1]]

Paths: 1 -> 3 -> 1 -> 1 -> 1  => 7
       1 -> 3 -> 5 -> 1 -> 1 => 11
       1 -> 3 -> 5 -> 2 -> 1 => 12
       1 -> 1 -> 5 -> 1 -> 1 => 9
       1 -> 1 -> 5 -> 2 -> 1 => 10
       1 -> 1 -> 4 -> 2 -> 1 => 9

*/

==============================================================================================

>> Minimum Path Sum <<

==============================================================================================

Problem: Find a path from the top left to bottom right which minimizes the sum of all numbers
along the path.

You may only move down or right.

Problem-solving strategy: bottom-up dynamic programming.

          GRID                                    CACHE
+-----------------------+               +-----------------------+
|       |       |       |               |       |       |       |
|   1   |   3   |   1   |               |   1   |   3   |   1   |
|_______|_______|_______|               |_______|_______|_______|
|       |       |       |               |       |       |       |
|   1   |   5   |   1   |               |   1   |   5   |   1   |
|_______|_______|_______|               |_______|_______|_______|
|       |       |       |               |       |       |       |
|   4   |   2   |   1   |               |   4   |   2   |   1   |
|_______|_______|_______|               |_______|_______|_______|

Iteratively build a map of the minimum path sums.

       -> If we're in square [0, 0], the minimum path will be the value at that coordinate.

       -> If we're in the first row or the first column, the minimum path will be the current
          value plus the sum of the previous values in that row/column.

       -> If we're not at the edge of the grid, we'll check the values saved in our cache for
          the squares above and to the left of the current value and select the minimum.
          We then add this minimum sum to the current value of our square and save it in our cache.

[0, 0] : 1
[0, 1] : 4
[0, 2] : 5
[1, 0] : 2
[1, 1] : 7
[1, 2] : 6
[2, 0] : 6
[2, 1] : 8
[2, 2] : 7

Once we reach the last square of the grid, we can return the minimum path sum!

==================================================================================================


package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	cache := make(map[[2]int]int)
	var previousSum int

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {

			if row == 0 {
				previousSum = cache[[2]int{0, column - 1}]
			} else if column == 0 {
				previousSum = cache[[2]int{row - 1, 0}]
			} else {
				previousSum = min(cache[[2]int{row - 1, column}], cache[[2]int{row, column - 1}])
			}
			cache[[2]int{row, column}] = previousSum + grid[row][column]
		}
		fmt.Println(cache)
	}
	return cache[[2]int{len(grid) - 1, len(grid[0]) - 1}]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
