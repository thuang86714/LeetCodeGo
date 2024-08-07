/*
There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.
Example 2:

Input: m = 3, n = 2
Output: 3
Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
1. Right -> Down -> Down
2. Down -> Down -> Right
3. Down -> Right -> Down
*/
package leetcodego
func uniquePaths(m int, n int) int {
    //credit to brianchang_tw
    //dp problem
    dp := make([][]int, m)
    temp := make([]int, m*n)
    for i := 0; i < m; i++{
        dp[i] = temp[i*n:(i+1)*n]
    }
    for i:=0; i<m;i++{
        for j:=0; j<n; j++{
            if i == 0 || j==0{
                dp[i][j]=1
            }else{
                dp[i][j] = dp[i-1][j] + dp[i][j-1]
            }

        }
    }
    return dp[m-1][n-1]


}