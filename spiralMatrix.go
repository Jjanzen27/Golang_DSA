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
*/

package main

import "fmt"

func spiralMatrix(matrix [][]int) (spiral []int) {
	direction := "r"
	level := 1
	i := 0
	j := 0

	for len(spiral) < 12 { // half as many times as the outer matrix
		fmt.Println(direction, j, i, spiral)
		if direction == "r" {
			spiral = append(spiral, matrix[j][i])
			if i >= len(matrix[0])-level {
				direction = "d"
			} else {
				i++
			}
		} else if direction == "d" {
			j++
			spiral = append(spiral, matrix[j][i])
			if j >= len(matrix)-level {
				direction = "l"
			}
		} else if direction == "l" {
			i--
			spiral = append(spiral, matrix[j][i])
			if i <= 0 {
				direction = "u"
			}
		} else if direction == "u" {
			j--
			spiral = append(spiral, matrix[j][i])
			if j <= 0 {
				direction = "r"
				level++
				j++
			}
		}
	}
	return spiral
}

func main() {
	fmt.Println(spiralMatrix([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))             // [1,2,3,6,9,8,7,4,5]
	fmt.Println(spiralMatrix([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}})) //[1,2,3,4,8,12,11,10,9,5,6,7]
}
