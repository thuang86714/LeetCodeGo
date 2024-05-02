/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.
Example1:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example2:
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
*/
package leetcodego
func setZeroes(matrix [][]int)  {
    isColzero := false
    isRowzero := false
    for i := 0; i < len(matrix); i++{
        if matrix[i][0] == 0{
            isColzero = true
            break
        }
    }
    for i := 0; i < len(matrix[0]); i++{
        if matrix[0][i] == 0{
            isRowzero = true
            break
        }
    }
    for i := 1; i < len(matrix); i++{
        for j := 1; j < len(matrix[0]); j++{
            if matrix[i][j]==0{
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }
    for i := 1; i < len(matrix); i++{
        for j := 1; j < len(matrix[0]); j++{
            if matrix[i][0] == 0 || matrix[0][j] == 0{
                matrix[i][j]= 0
            }
        }
    }
    if isColzero == true{
        for i:=0; i < len(matrix); i++{
            matrix[i][0]=0
        }
    }
    if isRowzero == true {
        for i:=0; i < len(matrix[0]); i++{
            matrix[0][i]= 0
        }
    }
}