package leetcodego

import(
	"sort"
)

type numFreqPair struct {
	num, freq int
}

func frequencySort(nums []int) []int {
	numFreqMap := make(map[int]int)
	numFreqPairSlice := []numFreqPair{}
	for _, num := range nums {
		numFreqMap[num]++
	}

	for curNum, curFreq := range numFreqMap {
		numFreqPairSlice = append(numFreqPairSlice, numFreqPair{
			num:  curNum,
			freq: curFreq,
		})
	}

	sort.Slice(numFreqPairSlice, func(i, j int) bool {
		if numFreqPairSlice[i].freq != numFreqPairSlice[j].freq {
			return numFreqPairSlice[i].freq < numFreqPairSlice[j].freq
		}
		return numFreqPairSlice[i].num > numFreqPairSlice[j].num
	})

	result := []int{}
	for _, pair := range numFreqPairSlice {
		for i := 0; i < pair.freq; i++ {
			result = append(result, pair.num)
		}
	}
	return result
}