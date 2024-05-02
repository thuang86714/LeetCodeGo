package leetcodego

func minimumTotal(triangle [][]int) int {
    //TC O(n), SC O(1)
    height := len(triangle)
    for i:= height - 2; i >= 0; i--{
        for j := range triangle[i]{
            triangle[i][j] += min(triangle[i + 1][j], triangle[i + 1][j + 1])
        }
    }
    return triangle[0][0]
}