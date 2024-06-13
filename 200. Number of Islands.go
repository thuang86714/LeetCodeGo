package leetcodego


func findIslandDFS(x, y int, grid [][]byte){
    grid[x][y] = '0'
    for _, dir := range dirs{
        newX, newY := x + dir[0], y + dir[1]
        if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == '1'{
            findIslandDFS(newX, newY, grid)
        } 
    }
}
func numIslands2(grid [][]byte) int {
    //DFS, TC O(n), SC O(1)
    islandCount := 0
    for i:= 0; i < len(grid); i++{
        for j := 0; j < len(grid[0]); j++{
            if grid[i][j] == '1'{
                islandCount++
                findIslandDFS(i, j, grid)
            }
        }
    }
    return islandCount
}