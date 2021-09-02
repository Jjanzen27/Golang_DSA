/* Given an m x n matrix, return all elements of the matrix in spiral order.



Example 1:


Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:


Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]


Constraints:

m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100

func spiralOrder(matrix [][]int) []int {
  result := make([]int, 0)
  if len(matrix) == 0 || len(matrix[0]) == 0 {
    return result
  }

  top := 0
  bottom := len(matrix) - 1
  left := 0
  right := len(matrix[0]) - 1
  size := len(matrix) * len(matrix[0])

  for len(result) < size {
    for i := left; i <= right && len(result) < size; i++ {
      result = append(result, matrix[top][i])
    }
    top++
    for i := top; i <= bottom && len(result) < size; i++ {
      result = append(result, matrix[i][right])
    }
    right--
    for i := right; i >= left && len(result) < size; i-- {
      result = append(result, matrix[bottom][i])
    }
    bottom--
    for i := bottom; i >= top && len(result) < size; i-- {
      result = append(result, matrix[i][left])
    }
    left++
  }
  return result
}

*/

package main

import "fmt"

func spiralMatrix(matrix [][]int) (spiral []int) {
	direction := "r"
	i := 0
	j := 0
	stepsAcross := len(matrix[0]) - 1
	stepsDown := len(matrix) - 1
	level := 0

	for level < (stepsAcross*stepsDown)/2 { // half as many times as the outer matrix
		fmt.Println(direction, j, i, spiral)
		if direction == "r" {
			spiral = append(spiral, matrix[j][i])
			if i >= stepsAcross-level {
				direction = "d"
				j++
			} else {
				i++
			}
		} else if direction == "d" {
			spiral = append(spiral, matrix[j][i])
			if j >= stepsDown-level {
				direction = "l"
				i--
			} else {
				j++
			}
		} else if direction == "l" {
			spiral = append(spiral, matrix[j][i])
			if i <= 0+level {
				direction = "u"
				j--
			} else {
				i--
			}
		} else if direction == "u" {
			level++
			spiral = append(spiral, matrix[j][i])
			if j <= 0+level {
				direction = "r"
				i++
				j++
			} else {
				j--
			}
		}
	}
	return spiral
}

func main() {
	// 	fmt.Println(spiralMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))             // [1,2,3,6,9,8,7,4,5]
	fmt.Println(spiralMatrix([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})) //[1,2,3,4,8,12,11,10,9,5,6,7]
}
