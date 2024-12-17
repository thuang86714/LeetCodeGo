package leetcodego

func kthDistinct(arr []string, k int) string {
	stringCount := make(map[string]int)

	for _, str := range arr {
		stringCount[str]++
	}

	distinctStr := []string{}
	for _, str := range arr {
		if stringCount[str] == 1 {
			distinctStr = append(distinctStr, str)
		}
	}

	var ans string
	for idx, str := range distinctStr {
		if idx == k-1 {
			ans = str
		}
	}
	return ans
}