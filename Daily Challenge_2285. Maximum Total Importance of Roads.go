package leetcodego

func maximumImportance(n int, roads [][]int) int64 {
	//TC O(n), SC O(n), counting sort
	degreeCountSlice := make([]int, n)
	for _, road := range roads {
		degreeCountSlice[road[0]]++
		degreeCountSlice[road[1]]++
	}
	freqCountSlice := make([]int, len(degreeCountSlice)+1)

	for _, degree := range degreeCountSlice {
		freqCountSlice[degree]++
	}

	var ans int64
	importance := n
	for i := len(freqCountSlice) - 1; i >= 0; i-- {
		for freqCountSlice[i] > 0 {
			ans += int64(i * importance)
			importance--
			freqCountSlice[i]--
		}
	}

	return ans
}