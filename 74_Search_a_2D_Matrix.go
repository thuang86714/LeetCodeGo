/*
You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

Example:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
*/
package leetcodego
func searchMatrix(matrix [][]int, target int) bool {
    rows, row, col := len(matrix), 0, len(matrix[0])-1
    for row < rows && col > -1{
        cur := matrix[row][col]
        if cur == target{
            return true
        }
        if target > cur {
            row++
        }else{
            col--
        }
    }
    return false

}

func searchMatrix2(matrix [][]int, target int) bool {
	n, left, right := len(matrix[0]), 0, len(matrix)*len(matrix[0])-1
	var middle, row, col, value int
	for left <= right{
		middle = (left+right)/2
		row = middle/n
		col = middle%n
		value = matrix[row][col]
		if value == target {
			return true
		}else if value > target{
			right = middle -1
		}else{
			left = middle +1
		}
	}
	return false
}