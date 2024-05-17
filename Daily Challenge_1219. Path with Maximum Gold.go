package leetcodego

//var dirs = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
func isValid(newX, newY int, grid [][]int)bool{
    if(newX < 0 || newX >= len(grid) || newY < 0 || newY >= len(grid[0])){
        return false
    }
    return grid[newX][newY] > 0
}
func collectGold(x, y int, grid [][]int)int{
    curGold, maxNextStep := grid[x][y], 0
    grid[x][y] = 0 
    for i := range dirs{
        newX := x + dirs[i][0]
        newY := y + dirs[i][1]
        if isValid(newX, newY, grid){
            maxNextStep = max(maxNextStep, collectGold(newX, newY, grid))
        }
    }
    grid[x][y] = curGold
    return curGold + maxNextStep
}
func getMaximumGold(grid [][]int) int {
    /*
    Time: O(k * 3 ^ k), where k is the number of cells with gold. We perform the analysis for k cells, and from each cell we can go in three directions.
    Memory: O(k) for the recursion.
    */
    ans := 0
    for i := range grid{
        for j := range grid[0]{
            ans = max(ans, collectGold(i, j, grid))
        }
    }
    return ans
}