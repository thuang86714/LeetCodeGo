/*
Given a positive integer n, generate an n x n matrix filled with elements from 1 to n2 in spiral order.
Example:
Input: n = 3
Output: [[1,2,3],[8,9,4],[7,6,5]]
Example 2:

Input: n = 1
Output: [[1]]*/
package leetcodego
func generateMatrix(n int) [][]int {
    //credit to oizg
    result := make([][]int, n)
    for i := 0; i < n; i++ { result[i] = make([]int, n) }
    dirs := [][]int{{0,1}, {1, 0}, {0, -1}, {-1, 0}}
    i, j, d := 0, 0, 0
    for k:=1; k <=n*n;k++{
        result[i][j]= k
        di, dj := dirs[d%4][0], dirs[d%4][1]
        //check if a) got out of bounds, or b) got previously filled cell.
        if i+di<0|| i+di>=n || j+dj<0 || j+dj>=n || result[i+di][j+dj] !=0{
            d++
            di, dj = dirs[d%4][0], dirs[d%4][1]
        }
        i, j = i+di, j+dj
    }
    return result
}