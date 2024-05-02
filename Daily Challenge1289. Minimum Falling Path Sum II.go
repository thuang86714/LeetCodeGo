package leetcodego
import(
	"math"
)
func minFallingPathSum(grid [][]int) int {
    //credit to votrubac, TC O(n^2), SC O(1)
    /*
    credit to taserface:
    If firstMinCur of two rows share the same column, 
    then take 
    min(firstMinCur + secondMinNext, secondMinCur + firstMinNext)
    */
    firstMinCur, secondMinCur, posCur, len := 0, 0, -1, len(grid)
    for i := 0; i < len; i++{
        firstMinNext, secondMinNext, posNext := math.MaxInt32, math.MaxInt32, -1
        for j := 0; j < len; j++{
            var curMemo int
            if j == posCur{
                curMemo = secondMinCur
            }else{
                curMemo = firstMinCur
            }
            if grid[i][j] + curMemo < firstMinNext{
                secondMinNext = firstMinNext
                firstMinNext = grid[i][j] + curMemo
                posNext = j
            }else{
                secondMinNext = min(secondMinNext, grid[i][j] + curMemo)
            }
        }
        firstMinCur, secondMinCur, posCur = firstMinNext, secondMinNext, posNext
    }
    return firstMinCur
}