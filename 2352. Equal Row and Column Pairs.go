package leetcodego

import "fmt"

// EqualPairs counts the number of equal pairs of rows and columns in a grid.
// It compares each row with each column for equality, considering the grid's size.
// Time Complexity: O(n^2), Space Complexity: O(n^2).
func EqualPairs(grid [][]int) int {
    gridSize := len(grid)
    transposedGrid := make([][]int, gridSize)
    for i := range transposedGrid {
        transposedGrid[i] = make([]int, gridSize)
    }

    // Transpose the grid
    for row, rowVals := range grid {
        for col := range rowVals {
            transposedGrid[col][row] = grid[row][col]
        }
    }

    equalPairsCount := 0
    // Compare each row with each column
    for row := 0; row < gridSize; row++ {
        for col := 0; col < gridSize; col++ {
            if equal(grid[row], transposedGrid[col]) {
                equalPairsCount++
            }
        }
    }

    return equalPairsCount
}

// Helper function to check if two slices are equal
func equal(slice1, slice2 []int) bool {
    for i, v := range slice1 {
        if v != slice2[i] {
            return false
        }
    }
    return true
}

func main() {
    grid := [][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }
    fmt.Println(EqualPairs(grid)) // Output: 0
}
