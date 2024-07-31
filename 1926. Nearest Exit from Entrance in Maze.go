package leetcodego
func nearestExit(maze [][]byte, entrance []int) int {
    dirs := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
    m, n := len(maze), len(maze[0])
	toVisit := make([][]int, 0)
	toVisit = append(toVisit, []int{entrance[0], entrance[1]})
	maze[entrance[0]][entrance[1]] = '+'
	ans := 0

	for len(toVisit) > 0 {
		curSize := len(toVisit)
		for i := 0; i < curSize; i++ {
			row, col := toVisit[0][0], toVisit[0][1]
			toVisit = toVisit[1:] // Dequeue

			if (row == 0 || col == 0 || row == m-1 || col == n-1) && !(row == entrance[0] && col == entrance[1]) {
				// Check if we are on the border but not at the entrance
				return ans
			}

			for _, dir := range dirs {
				newRow, newCol := row+dir[0], col+dir[1]
				if newRow >= 0 && newCol >= 0 && newRow < m && newCol < n && maze[newRow][newCol] == '.' {
					toVisit = append(toVisit, []int{newRow, newCol}) // Enqueue
					maze[newRow][newCol] = '+'
				}
			}
		}
		ans++
	}

	return -1
}