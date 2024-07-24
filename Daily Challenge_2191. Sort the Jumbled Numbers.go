package leetcodego
import(
	"sort"
)
type numIdxCombo struct {
	oldNum, newNum, idx int
}

func sortJumbled(mapping []int, nums []int) []int {
	//TC O(n∗d+nlogn), SC O(nlogn)
	numIdxSlice := []numIdxCombo{}
	for curIdx, num := range nums {
		cur, next, digit := num, 0, 1
		if cur == 0 {
			numIdxSlice = append(numIdxSlice, numIdxCombo{
				oldNum: num,
				newNum: mapping[0],
				idx:    curIdx,
			})
			continue
		}

		for cur > 0 {
			next = mapping[cur%10]*digit + next
			cur /= 10
			digit *= 10
		}
		numIdxSlice = append(numIdxSlice, numIdxCombo{
			oldNum: num,
			newNum: next,
			idx:    curIdx,
		})
	}
	sort.Slice(numIdxSlice, func(i, j int) bool {
		if numIdxSlice[i].newNum == numIdxSlice[j].newNum {
			return numIdxSlice[i].idx < numIdxSlice[j].idx
		}
		return numIdxSlice[i].newNum < numIdxSlice[j].newNum
	})

	result := []int{}
	for _, combo := range numIdxSlice {
		result = append(result, combo.oldNum)
	}

	return result
}