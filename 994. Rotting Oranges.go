func orangesRotting(grid [][]int) int {
    dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    m, n := len(grid), len(grid[0])
	freshOrange, ans := 0, -1
	toVisit := make([][2]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				freshOrange++
			} else if grid[i][j] == 2 {
				toVisit = append(toVisit, [2]int{i, j})
			}
		}
	}

	for len(toVisit) > 0 {
		curSize := len(toVisit)
		for i := 0; i < curSize; i++ {
			row, col := toVisit[0][0], toVisit[0][1]
			toVisit = toVisit[1:]
			for _, dir := range dirs {
				newRow, newCol := row+dir[0], col+dir[1]
				if newRow >= 0 && newCol >= 0 && newRow < m && newCol < n && grid[newRow][newCol] == 1 {
					grid[newRow][newCol] = 2
					toVisit = append(toVisit, [2]int{newRow, newCol})
					freshOrange--
				}
			}
		}
		ans++
	}

	if freshOrange > 0 {
		return -1
	}
	if ans == -1 {
		return 0
	}
	return ans
}