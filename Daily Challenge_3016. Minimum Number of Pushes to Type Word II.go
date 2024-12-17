package leetcodego

import(
	"sort"
)

type charCount struct {
	alphabet rune
	count    int
}

func minimumPushes(word string) int {
	count := 0
	runeCountMap := make(map[rune]int)
	for _, char := range word {
		runeCountMap[char]++
	}
	charCountSlice := make([]charCount, 0)

	for k, v := range runeCountMap {
		cur := charCount{
			alphabet: k,
			count:    v,
		}
		charCountSlice = append(charCountSlice, cur)
	}

	sort.Slice(charCountSlice, func(i, j int) bool {
		return charCountSlice[i].count > charCountSlice[j].count
	})

	for i := 0; i < len(charCountSlice); i++ {
		if i < 8 {
			count += charCountSlice[i].count
			continue
		}

		if i < 16 {
			count += charCountSlice[i].count * 2
			continue
		}

		if i < 24 {
			count += charCountSlice[i].count * 3
			continue
		}

		count += charCountSlice[i].count * 4
	}
	return count
}