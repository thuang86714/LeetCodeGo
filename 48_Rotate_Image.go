/*
You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.
Example:
Input: matrix = [[1,2,3],[4,5,6],[7,
*/
func rotate(matrix [][]int)  {
	//credit to pixalquarks
    n := len(matrix)
    for i := 0; i < n; i++{
        for j:= i; j < n; j++{
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
        for j, k := 0, n-1; j<k; j, k = j+1, k-1{
            matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
        }
    }
}