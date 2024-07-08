package leetcodego

func findTheWinner(n int, k int) int {
	/*
	   TC O(n), SC O(1)
	   credit to hritesh_bhardwaj
	   Intution : If I have the winner for 'n-1' and 'k'.
	   I can find the winner for 'n' and 'k' by moving on to the next kth
	   person (i.e. f(n) = f(n-1)+k).
	*/
	ans := 0
	for i := 1; i <= n; i++ {
		ans = (ans + k) % i
	}
	return ans + 1
}

/*
func findTheWinner(n int, k int) int {
	circleSlice := []int{}
	for i := 1; i <= n; i++ {
		circleSlice = append(circleSlice, i)
	}

	for len(circleSlice) > 1 {
		curK, lastNode := k, 0
		for curK > 0 {
			if lastNode != 0 {
				circleSlice = append(circleSlice, lastNode)
			}
			lastNode = circleSlice[0]
			circleSlice = circleSlice[1:]
			curK--
		}
	}
	return circleSlice[0]
}
*/