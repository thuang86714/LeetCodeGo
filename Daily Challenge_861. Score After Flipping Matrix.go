package leetcodego

//the function's cognitive complexity is 33, which is way too high, 
//solution to this is coming up two extra function: iterateByRow and iterateByCol, both return int(rowSum and colSum)
func matrixScore(grid [][]int) int {
    //TC O(m*n), SC O(1)
    //the key is to get as more 1 as we can
    //iterate by row to see if this row needs to be fliped
	
    for i := range grid{
        rowSum := 0
        for j := range grid[0]{
            //every element in a row is weighted
            rowSum *= 2
            if grid[i][j] == 1{
                rowSum += 1
            }else{
                rowSum -= 1
            }
        }
        //early return
        if rowSum >= 0{
            continue
        }
        for j := range grid[0]{
            if grid[i][j] == 1{
                grid[i][j] = 0
            }else{
                grid[i][j] = 1
            }
        }
    }
    
    
    //iterate by column to see if this row needs to be fliped
    for j := range grid[0]{
        colSum := 0
        for i:= range grid{
            if grid[i][j] == 1{
                colSum += 1
            }else{
                colSum -= 1
            }
        }
        //early return
        if colSum >= 0{
            continue
        }
        for i := range grid{
            if grid[i][j] == 1{
                grid[i][j] = 0
            }else{
                grid[i][j] = 1
            }
        }
    }
    
    ans := 0
    for i := range grid{
        rowSum := 0
        for j := range grid[0]{
            rowSum *= 2
            rowSum += grid[i][j]
        }
        ans += rowSum
    }
    return ans
}