package leetcodego

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func findThisFarmland(x, y int, land [][]int, curCoordinates []int, ans *[][]int) {
    land[x][y] = 0  // Mark the cell as visited by setting it to 0
    for _, dir := range dirs {
        newX := x + dir[0]
        newY := y + dir[1]
        if newX >= 0 && newX < len(land) && newY >= 0 && newY < len(land[0]) && land[newX][newY] == 1 {
            if newX > curCoordinates[2] {
                curCoordinates[2] = newX
            }
            if newY > curCoordinates[3] {
                curCoordinates[3] = newY
            }
            findThisFarmland(newX, newY, land, curCoordinates, ans)
        }
    }
    // Append current farmland boundaries if this is the root call for a new farmland
    if x == curCoordinates[0] && y == curCoordinates[1] {
        *ans = append(*ans, curCoordinates)
    }
}

func findFarmland(land [][]int) [][]int {
    ans := [][]int{}
    for i := 0; i < len(land); i++ {
        for j := 0; j < len(land[0]); j++ {
            if land[i][j] == 1 {
                temp := []int{i, j, i, j}
                findThisFarmland(i, j, land, temp, &ans)
            }
        }
    }
    return ans
}
