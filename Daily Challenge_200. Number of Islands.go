package leetcodego

var dirs = [][]int{{1,0}, {0, 1}, {-1, 0}, {0, -1}}
func findThisIsland(x, y int, grid [][]byte){
    //dfs
    if grid[x][y] == '0'{
        return
    }

    grid[x][y] = '0'

    for _,dir := range dirs{
        newX := x + dir[0]
        newY := y + dir[1]
        if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]){
            findThisIsland(newX, newY, grid)
        }
    }
}
func numIslands(grid [][]byte) int {
    // TC O(m*n), SC O(m*n)
    islandCount := 0;
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[0]); j++{
            if grid[i][j] == '1'{
                islandCount++
                findThisIsland(i, j, grid)
            }
        }
    }
    return islandCount;
}

/*
var dirs = [][]int{{1,0}, {0, 1}, {-1, 0}, {0, -1}}
type pair struct{
    X int
    Y int
}
func findThisIsland(x, y int, grid [][]byte){
    //bfs
    toVisitSlice := []pair{}
    toVisitSlice = append(toVisitSlice, pair{X: x, Y: y})
    grid[x][y] = '0'
    for len(toVisitSlice) != 0{
        curX, curY := toVisitSlice[0].X, toVisitSlice[0].Y
        toVisitSlice = toVisitSlice[1:]

        for _, dir := range dirs{
            newX := curX + dir[0]
            newY := curY + dir[1]
            if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && grid[newX][newY] == '1'{
                toVisitSlice = append(toVisitSlice, pair{X: newX, Y: newY})
                grid[newX][newY] = '0'
            }
        }
    }
}
func numIslands(grid [][]byte) int {
    // TC O(m*n), SC O(m*n)
    islandCount := 0;
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[0]); j++{
            if grid[i][j] == '1'{
                islandCount++
                findThisIsland(i, j, grid)
            }
        }
    }
    return islandCount;
}
*/