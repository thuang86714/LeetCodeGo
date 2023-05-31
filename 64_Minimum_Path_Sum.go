/*Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example1:
Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
Output: 7
Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
Example 2:

Input: grid = [[1,2,3],[4,5,6]]
Output: 12*/

func minPathSum(grid [][]int) int {
    //can't go greedy, local minimum doesn't lead to total minimum
    //TC O(n), SC O(1)
    m := len(grid)
    n := len(grid[0])
    for i:=1; i < m; i++{
        grid[i][0] += grid[i-1][0]
    }
    for i:=1; i < n; i++{
        grid[0][i] += grid[0][i-1]
    } 
    for i:=1; i < m; i++{
        for j:=1; j < n ; j++{
            grid[i][j] += min(grid[i-1][j], grid[i][j-1])
        }
    }
    return grid[m-1][n-1]
}

func min (x, y int) int{
    if x >= y{
        return y
    }
    return x
}