package leetcodego

func dailyTemperatures(temperatures []int) []int {
	//credit to riemeltm, TC O(n), SC O(n)
    indexSlice, result := []int{}, make([]int, len(temperatures))

    for idx := len(temperatures) - 1; idx >= 0; idx--{
        // Pop elements from indexSlice while the current temperature is greater or equal to
        // the temperature at the index on top of the indexSlice stack
        for len(indexSlice) > 0 && temperatures[indexSlice[len(indexSlice) -1]] <= temperatures[idx]{
            indexSlice = indexSlice[:len(indexSlice) - 1]
        }

        // If indexSlice is not empty, calculate the days until a warmer temperature
        if len(indexSlice) > 0{
            result[idx] = indexSlice[len(indexSlice) - 1] - idx
        }

        // Push the current index onto indexSlice
        indexSlice = append(indexSlice, idx) 
    }
    return result
}