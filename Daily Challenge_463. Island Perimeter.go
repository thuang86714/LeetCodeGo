package leetcodego

func islandPerimeter(grid [][]int) int {
    //credit to lamXiaoBai, TC O(m*n), SC O(1)
    cellCount, repeat := 0, 0
    for i := 0; i < len(grid); i++{
        for j := 0; j < len(grid[0]); j++{
            if grid[i][j] == 0{
                continue
            }
			cellCount++
			if i > 0 && grid[i - 1][j] == 1{
				repeat++
			}
			if j > 0 && grid[i][j - 1] == 1{
				repeat++
			}
        }
    }
    return 4*cellCount - 2 *repeat
}