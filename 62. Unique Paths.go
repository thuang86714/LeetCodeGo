package leetcodego

func uniquePaths2(m int, n int) int {
	memoSlice := make([]int, n)
	for i := 0; i < n; i++ {
		memoSlice[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			memoSlice[j] += memoSlice[j-1]
		}
	}
	return memoSlice[n-1]
}
