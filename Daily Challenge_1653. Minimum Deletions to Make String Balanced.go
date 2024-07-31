package leetcodego

func minimumDeletions(s string) int {
    //credit to neetcode and arishta, TC O(n), SC O(1)
	aCountRight, bCountLeft, result := 0, 0, len(s)

	for _, r := range s {
		if r == 'a' {
			aCountRight++
		}
	}

	for _, r := range s {
		if r == 'a' {
			aCountRight--
		}
		result = min(result, aCountRight+bCountLeft)
		if r == 'b' {
			bCountLeft++
		}
	}
	return result
}