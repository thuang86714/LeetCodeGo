package leetcodego


func lengthOfLastWord2(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && ans != 0 {
			break
		}
		if s[i] == ' ' {
			continue
		}
		ans++
	}
	return ans
}