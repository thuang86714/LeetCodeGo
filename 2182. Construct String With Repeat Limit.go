package leetcodego

func repeatLimitedString(s string, repeatLimit int) string {
	charCountSlice := make([]int, 26)
	//do char count
	for _, char := range s {
		charCountSlice[char-'a']++
	}
	//iterate char count array
	ans := []byte{}
	found, count, last := true, 0, -1
	for found {
		found = false
		for i := 25; i >= 0 && !found; i-- {
			if charCountSlice[i] > 0 && (count < repeatLimit || last != i) {
				ans = append(ans, byte('a'+i))
				count++
				charCountSlice[i]--
				if last != i {
					count = 1
				}
				last = i
				found = true
			}
		}
	}
	return string(ans)
}