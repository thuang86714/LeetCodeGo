/*
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
Example 2:

Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
*/
func isValidSudoku(board [][]byte) bool {
    //credit to sandeepmanne
    row := [9][9]bool{}//first [] represent numbering of row, second represent val
    col := [9][9]bool{}
    grid := [9][9]bool{}
    for i := 0; i < 9; i++{
        for j := 0; j< 9; j++{
            val, err := strconv.Atoi(string(board[i][j]))
            if err != nil{//which means it's "."
                continue
            }
            val--//number range 1-9, but row, col, grid are 0-indexed
            gridIdx := (i/3)*3 + j/3//create an unique numbering for grid, range from 0`~8`
            if row[i][val] || col[j][val] || grid[gridIdx][val]{
                return false
            }
            row[i][val] = true
            col[j][val] = true    
            grid[gridIdx][val] = true
        }
    }
    return true
}