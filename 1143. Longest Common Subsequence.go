package leetcodego

func longestCommonSubsequence(text1 string, text2 string) int {
	curMemo := make([][]int, 1001)
	for i := range curMemo {
		curMemo[i] = make([]int, 1001)
	}

	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				curMemo[i+1][j+1] = curMemo[i][j] + 1
			} else {
				curMemo[i+1][j+1] = max(curMemo[i][j+1], curMemo[i+1][j])
			}
		}
	}
	return curMemo[len(text1)][len(text2)]
}