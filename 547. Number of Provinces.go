package leetcodego

func findAllConnected(isConnected [][]int, visitedSlice []int, idx int){
	for j := 0; j < len(isConnected); j++{
		if isConnected[idx][j] == 1 && visitedSlice[j] == 0{
			visitedSlice[j] = 1
			findAllConnected(isConnected, visitedSlice, j)
		}
	}
}
func findCircleNum(isConnected [][]int) int {
	//credit to vindo23, TC O(n^2), SC O(n)
visitedSlice := make([]int, len(isConnected))

provinceCount := 0
for i := 0; i < len(isConnected); i++{
	if visitedSlice[i] == 0{
		visitedSlice[i] = 1;
		findAllConnected(isConnected, visitedSlice, i)
		provinceCount++
	}
}
return provinceCount
}