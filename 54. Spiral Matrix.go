func spiralOrder(matrix [][]int) []int {
    //TC O(m*n), SC O(1)
    left, right, top, bottom := 0, len(matrix[0]) - 1, 0, len(matrix) - 1
    res := []int{}
    for left <= right && top <= bottom{
        for j := left; j <= right && top <= bottom; j++{
            res = append(res, matrix[top][j])
        }
        top++

        for i := top; i <= bottom && left <= right; i++{
            res = append(res, matrix[i][right])
        }
        right--

        for j := right; left <= j && top <= bottom; j--{
            res = append(res, matrix[bottom][j])
        }
        bottom--

        for i := bottom; top <= i && left <= right; i--{
            res = append(res, matrix[i][left])
        }
        left++
        
    }
    return res
}