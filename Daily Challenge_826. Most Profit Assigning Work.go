package leetcodego
import(
	"sort"
)

type Reward struct{
    diff int
    prof int
}

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    //TC O(nlogn+mlogm), where m is the len(difficulty), SC O(n + m)
	sum := 0

	// Sort the worker array
	sort.Ints(worker)

	// Create the rewardSlice
	rewardSlice := make([]Reward, len(difficulty))
	for idx := range difficulty {
		rewardSlice[idx] = Reward{
			diff: difficulty[idx],
			prof: profit[idx],
		}
	}

	// Sort the rewardSlice by difficulty, and by profit in case of ties
	sort.Slice(rewardSlice, func(i, j int) bool {
		if rewardSlice[i].diff != rewardSlice[j].diff {
			return rewardSlice[i].diff < rewardSlice[j].diff
		}
		return rewardSlice[i].prof > rewardSlice[j].prof
	})

	// Maintain a list of maximum profits up to the current difficulty
	maxProfits := make([]int, len(rewardSlice))
	maxProfits[0] = rewardSlice[0].prof
	for i := 1; i < len(rewardSlice); i++ {
		maxProfits[i] = max(maxProfits[i-1], rewardSlice[i].prof)
	}

	idx := 0
	for _, ability := range worker {
		// Use binary search to find the maximum difficulty the worker can handle
		for idx < len(rewardSlice) && ability >= rewardSlice[idx].diff {
			idx++
		}
		if idx > 0 {
			sum += maxProfits[idx-1]
		}
	}

	return sum
}