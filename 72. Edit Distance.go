package leetcodego

func minDistance(word1 string, word2 string) int {
    //credit to jianchao-li, TC O(m*n), SC O(m*n)
    m, n := len(word1), len(word2)
    memoSlice := make([][]int, m + 1)
    for idx := range memoSlice{
        memoSlice[idx] = make([]int, n + 1)
    }

    for i := 1; i <= m ; i++{
        memoSlice[i][0] = i
    }

    for j := 1; j <= n; j++{
        memoSlice[0][j] = j
    }

    for i := 1; i <= m; i++{
        for j := 1; j <= n; j++{
            if word1[i - 1] == word2[j - 1]{
                memoSlice[i][j] = memoSlice[i - 1][j - 1]
            }else{
                memoSlice[i][j] = min(memoSlice[i - 1][j - 1], min(memoSlice[i - 1][j], memoSlice[i][j - 1])) + 1
            }
        }
    }
    return memoSlice[m][n]
}