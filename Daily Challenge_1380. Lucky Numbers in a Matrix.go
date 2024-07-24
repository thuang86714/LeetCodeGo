package leetcodego

func luckyNumbers(matrix [][]int) []int {
	//TC O(m*n), SC O(m+n)
	m, n := len(matrix), len(matrix[0])
	maxCol, minRow, result := matrix[0], make([]int, m), []int{}

	for i := 0; i < m; i++ {
		curMin := 1000000
		for j := 0; j < n; j++ {
			maxCol[j] = max(maxCol[j], matrix[i][j])
			curMin = min(curMin, matrix[i][j])
		}
		minRow[i] = curMin
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if maxCol[j] == minRow[i] {
				result = append(result, maxCol[j])
			}
		}
	}
	return result
}