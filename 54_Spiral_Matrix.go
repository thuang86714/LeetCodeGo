/*
Given an m x n matrix, return all elements of the matrix in spiral order.
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
*/
package leetcodego
func spiralOrder(matrix [][]int) []int {
    n := len(matrix)
    m := len(matrix[0])
    result := []int{}
    top, bottom, left, right := 0, n-1, 0, m-1
    for{
        //go right
        for i:=left; i <=right; i++{
            result = append(result, matrix[top][i])
        }
        top++
        if top > bottom{
            break
        }

        //go down
        for i:=top; i <= bottom; i++{
            result = append(result, matrix[i][right])
        }
        right--
        if left > right{
            break
        }

        //go left
        for i:= right; i >=left; i--{
            result = append(result, matrix[bottom][i])

        }
        bottom--
        if bottom < top{
            break
        }

        //go up
        for i:= bottom; i >= top; i--{
            result =append(result, matrix[i][left])
        }
        left++
        if left > right{
            break
        }
    }
    return result
}