package leetcodego
import(
	"sort"
)
func sortPeople(names []string, heights []int) []string {
	heightNameMap := make(map[int]string)
	for idx := range heights {
		heightNameMap[heights[idx]] = names[idx]
	}

	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})

	result := []string{}
	for _, height := range heights {
		result = append(result, heightNameMap[height])
	}
	return result
}