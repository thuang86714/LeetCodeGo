package leetcodego

func largestLocal(grid [][]int) [][]int {
    //TC O(n^2), SC O(1)
    n := len(grid)
    result := make([][]int, n - 2)
    for idx := range result{
        result[idx] = make([]int, n - 2)
    }

    for i := 0; i < n - 2; i++{
        for j := 0; j < n - 2; j++{
            for I := i; I < i + 3; I++{
                for J := j; J < j + 3; J++{
                    result[i][j] = max(result[i][j], grid[I][J])
                }
            }
        }
    }
    return result
}