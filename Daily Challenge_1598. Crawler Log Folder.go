package leetcodego

func minOperations2(logs []string) int {
	ans := 0
	for _, log := range logs {
		if log == "../" {
			if ans > 0 {
				ans--
			}
		} else if log == "./" {
			continue
		} else {
			ans++
		}
	}
	if ans < 0 {
		return 0
	}
	return ans
}