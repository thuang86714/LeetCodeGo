/*
You are given an m x n integer array grid. There is a robot initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

An obstacle and space are marked as 1 or 0 respectively in grid. A path that the robot takes cannot include any square that is an obstacle.

Return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The testcases are generated so that the answer will be less than or equal to 2 * 109.
Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
Output: 2
Explanation: There is one obstacle in the middle of the 3x3 grid above.
There are two ways to reach the bottom-right corner:
1. Right -> Right -> Down -> Down
2. Down -> Down -> Right -> Right
*/
package leetcodego
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	//credit to brianchiang_tw
    m := len(obstacleGrid)
    n := len(obstacleGrid[0])
    var allowToVisit func (x, y int) int//in go, func could be a var as well
    
    allowToVisit = func(x, y int) int{//this helper func results in SC O(1)
        // 1 : allow to visit
        // 0 : can not visit due to obstacle
        return 1 - obstacleGrid[y][x]
    }
    if m*n == 0 || allowToVisit(0, 0) == 0{
        return 0
    }
    obstacleGrid[0][0] = 1

    for i := 1 ; i < m; i++{
        obstacleGrid[i][0]= obstacleGrid[i-1][0] * allowToVisit(0, i)
    }

    for j := 1; j < n; j++{
        obstacleGrid[0][j] = obstacleGrid[0][j-1] * allowToVisit(j, 0)
    }

    for i := 1; i < m ; i++{
        for j := 1; j < n; j++{
            obstacleGrid[i][j] = (obstacleGrid[i-1][j] + obstacleGrid[i][j-1]) *allowToVisit(j, i)
        }
    }
    return obstacleGrid[m-1][n-1]
}