package leetcodego

func averageWaitingTime(customers [][]int) float64 {
    //TC O(n), SC O(1)
	customerCount := len(customers)
	timeSum := 0
	startingTime, lastFinishTime := 0, customers[0][0]
	for _, order := range customers {
		startingTime = max(lastFinishTime, order[0])
		timeSum += (startingTime + order[1] - order[0])
		lastFinishTime = startingTime + order[1]
	}
	return float64(timeSum) / float64(customerCount)
}