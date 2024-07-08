package leetcodego
import(
	"sort"
)
func longestCommonPrefix2(strs []string) string {
	sort.Strings(strs)
	ans := 0
	frontStr, backStr := strs[0], strs[len(strs)-1]
	for i := 0; i < min(len(frontStr), len(backStr)); i++ {
		if frontStr[i] != backStr[i] {
			break
		}
		ans++
	}
	return frontStr[:ans]
}